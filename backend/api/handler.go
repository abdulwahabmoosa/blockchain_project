package api

import (
	"backend/auth"
	"backend/blockchain"
	"backend/blockchain/worker"
	"backend/db"
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

// Request payloads
type CreatePropertyRequest struct {
	OwnerAddress string `json:"owner_address"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	DataHash     string `json:"data_hash"` // IPFS Hash
	Valuation    int64  `json:"valuation"`
	TokenSupply  int64  `json:"token_supply"`
}

type DistributeRevenueRequest struct {
	TokenAddress      string `json:"token_address"`
	StablecoinAddress string `json:"stablecoin_address"`
	Amount            int64  `json:"amount"`
}

type RequestHandler struct {
	db    *db.Database
	chain *blockchain.ChainService
}

func NewRequestHandler(db *db.Database, chain *blockchain.ChainService) *RequestHandler {
	worker.StartListeners(chain, db)
	return &RequestHandler{db, chain}
}

func (handler *RequestHandler) Start() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Post("/login", handler.Login)
	r.Post("/register", handler.RegisterUser)

	// Protected Routes
	r.Group(func(r chi.Router) {
		r.Use(auth.Middleware)

		// Property Routes
		r.Get("/properties", handler.GetProperties)
		r.Get("/properties/{id}", handler.GetProperty)
		r.Post("/properties", handler.CreateProperty) // Triggers Blockchain

		// Revenue Routes
		r.Post("/revenue/distribute", handler.DistributeRevenue) // Triggers Blockchain
	})

	http.ListenAndServe(":3000", r)
}

func (handler *RequestHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds auth.LoginCredentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	user, err := handler.db.GetUserByCredentials(creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateToken(user.ID, user.Email)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	ctx := context.WithValue(r.Context(), auth.TokenKey, token)
	token, ok := ctx.Value(auth.TokenKey).(string)
	if !ok {
		status := http.StatusUnprocessableEntity
		http.Error(w, http.StatusText(status), status)
		return
	}

	render.JSON(w, r, map[string]string{"token": token})
}

func (handler *RequestHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var userDetails auth.RegisterUserPayload
	if err := json.NewDecoder(r.Body).Decode(&userDetails); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	exists, err := handler.db.UserExists(userDetails.Email)
	if exists {
		http.Error(w, "User Exists", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	err = handler.db.CreateUser(userDetails)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
}

func (handler *RequestHandler) GetProperties(w http.ResponseWriter, r *http.Request) {
	props, err := handler.db.GetAllProperties()
	if err != nil {
		http.Error(w, "Failed to fetch properties", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, props)
}

func (handler *RequestHandler) GetProperty(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	prop, err := handler.db.GetPropertyByID(id)
	if err != nil {
		http.Error(w, "Property not found", http.StatusNotFound)
		return
	}
	render.JSON(w, r, prop)
}

// Updated CreateProperty (Non-Blocking)
func (handler *RequestHandler) CreateProperty(w http.ResponseWriter, r *http.Request) {
	var req CreatePropertyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	// 1. Submit Transaction
	tx, err := handler.chain.CreateProperty(req.OwnerAddress, req.Name, req.Symbol, req.DataHash, req.Valuation, req.TokenSupply)
	if err != nil {
		http.Error(w, "Blockchain Submission Failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 2. Return Pending Response Immediately
	// The Worker will handle the DB creation when the event is emitted.
	render.JSON(w, r, map[string]string{
		"status":  "pending",
		"message": "Transaction submitted. Waiting for blockchain confirmation.",
		"tx_hash": tx.Hash().Hex(),
	})
}

// Updated DistributeRevenue (Non-Blocking)
func (handler *RequestHandler) DistributeRevenue(w http.ResponseWriter, r *http.Request) {
	var req DistributeRevenueRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	// 1. Submit Transaction
	tx, err := handler.chain.DistributeRevenue(req.TokenAddress, req.StablecoinAddress, req.Amount)
	if err != nil {
		http.Error(w, "Blockchain Submission Failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 2. Return Pending Response Immediately
	render.JSON(w, r, map[string]string{
		"status":  "pending",
		"message": "Revenue deposit submitted. Waiting for blockchain confirmation.",
		"tx_hash": tx.Hash().Hex(),
	})
}
