package api

import (
	"backend/auth"
	"backend/blockchain"
	"backend/blockchain/worker"
	"backend/db"
	"backend/db/models"
	"backend/ipfs"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware" // Ensure you ran: go get github.com/go-chi/cors
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10" // Ensure you ran: go get github.com/go-playground/validator/v10
)

// Initialize Validator
var validate = validator.New()

// --- Request Payloads with Validation Tags ---

type CreatePropertyRequest struct {
	OwnerAddress string `json:"owner_address" validate:"required,eth_addr"`
	Name         string `json:"name" validate:"required,min=3,max=100"`
	Symbol       string `json:"symbol" validate:"required,alphanum,len=3|len=4"` // 3 or 4 chars, e.g. "PROP"
	DataHash     string `json:"data_hash" validate:"required,printascii"`        // IPFS hash usually ascii
	Valuation    int64  `json:"valuation" validate:"required,gt=0"`              // Must be positive
	TokenSupply  int64  `json:"token_supply" validate:"required,gt=0"`
}

type DistributeRevenueRequest struct {
	TokenAddress      string `json:"token_address" validate:"required,eth_addr"`
	StablecoinAddress string `json:"stablecoin_address" validate:"required,eth_addr"`
	Amount            int64  `json:"amount" validate:"required,gt=0"`
}

type ApproveUserRequest struct {
	WalletAddress string `json:"wallet_address" validate:"required,eth_addr"`
}

type RequestHandler struct {
	db    *db.Database
	chain *blockchain.ChainService
}

func NewRequestHandler(db *db.Database, chain *blockchain.ChainService) *RequestHandler {
	// Start blockchain event listeners (optional - won't crash if subscriptions fail)
	if chain != nil {
		log.Printf("🔄 Starting blockchain event listeners...")
		worker.StartListeners(chain, db)
		log.Printf("✅ Blockchain listeners started (may show warnings if RPC doesn't support subscriptions)")
	}
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

	r.Group(func(r chi.Router) {
		r.Use(auth.Middleware)

		// Property Routes
		r.Get("/properties", handler.GetProperties)
		r.Get("/properties/{id}", handler.GetProperty)

		r.Post("/upload", handler.UploadMetadata)

		// Admin Routes
		r.Group(func(r chi.Router) {
			r.Use(handler.AdminMiddleware)
			r.Get("/users", handler.GetUsers)
			r.Post("/approve-user", handler.ApproveUser)
			r.Post("/properties", handler.CreateProperty)
			r.Post("/revenue/distribute", handler.DistributeRevenue)
		})
	})

	srv := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	// Start Server in a Goroutine so it doesn't block
	go func() {
		log.Println("🚀 Server starting on port 3000...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	// catch SIGINT (Ctrl+C) and SIGTERM (Docker stop)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("⚠️  Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("✅ Server exiting")
}

func (handler *RequestHandler) Login(w http.ResponseWriter, r *http.Request) {
	log.Printf("🔐 Login attempt started")

	var creds auth.LoginCredentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		log.Printf("❌ Login failed: JSON decode error: %v", err)
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	log.Printf("📋 Login credentials received: email=%s", creds.Email)

	// VALIDATION
	if err := validate.Struct(creds); err != nil {
		log.Printf("❌ Login failed: Validation error: %v", err)
		http.Error(w, "Validation Error: "+err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("✅ Login validation passed")

	user, err := handler.db.GetUserByCredentials(creds)
	if err != nil {
		log.Printf("❌ Login failed: Database error: %v", err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	log.Printf("✅ User found in database: ID=%s, Email=%s, Role=%s", user.ID, user.Email, user.Role)

	token, err := auth.GenerateToken(user.ID, user.Email)
	if err != nil {
		log.Printf("❌ Login failed: Token generation error: %v", err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	log.Printf("✅ Login token generated successfully")

	render.JSON(w, r, map[string]any{
		"token": token,
		"user":  user,
	})
}

func (handler *RequestHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var userDetails auth.RegisterUserPayload
	if err := json.NewDecoder(r.Body).Decode(&userDetails); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	// VALIDATION
	if err := validate.Struct(userDetails); err != nil {
		http.Error(w, "Validation Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	exists, err := handler.db.UserExists(userDetails.Email)
	if exists {
		http.Error(w, "User Exists", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = handler.db.CreateUser(userDetails)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]string{"message": "User registered successfully"})
}

func (handler *RequestHandler) CreateProperty(w http.ResponseWriter, r *http.Request) {
	var req CreatePropertyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	// VALIDATION
	if err := validate.Struct(req); err != nil {
		http.Error(w, "Validation Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	tx, err := handler.chain.CreateProperty(req.OwnerAddress, req.Name, req.Symbol, req.DataHash, req.Valuation, req.TokenSupply)
	if err != nil {
		http.Error(w, "Blockchain Submission Failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]string{
		"status":  "pending",
		"message": "Transaction submitted. Waiting for blockchain confirmation.",
		"tx_hash": tx.Hash().Hex(),
	})
}

func (handler *RequestHandler) DistributeRevenue(w http.ResponseWriter, r *http.Request) {
	var req DistributeRevenueRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	// VALIDATION
	if err := validate.Struct(req); err != nil {
		http.Error(w, "Validation Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	tx, err := handler.chain.DistributeRevenue(req.TokenAddress, req.StablecoinAddress, req.Amount)
	if err != nil {
		http.Error(w, "Blockchain Submission Failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]string{
		"status":  "pending",
		"message": "Revenue deposit submitted. Waiting for blockchain confirmation.",
		"tx_hash": tx.Hash().Hex(),
	})
}

func (handler *RequestHandler) ApproveUser(w http.ResponseWriter, r *http.Request) {
	// Middleware has already checked Admin role

	var req ApproveUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("❌ Failed to decode request: %v", err)
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	// VALIDATION
	if err := validate.Struct(req); err != nil {
		log.Printf("❌ Validation failed: %v", err)
		http.Error(w, "Validation Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Check if chain service is available
	if handler.chain == nil {
		log.Printf("❌ Chain service not initialized - blockchain functionality disabled")
		http.Error(w, "Blockchain service not available - check environment configuration", http.StatusServiceUnavailable)
		return
	}

	log.Printf("🔄 Approving user on blockchain: %s", req.WalletAddress)
	tx, err := handler.chain.ApproveUser(req.WalletAddress)
	if err != nil {
		log.Printf("❌ Blockchain approval failed for %s: %v", req.WalletAddress, err)
		http.Error(w, "Blockchain Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("✅ Blockchain approval confirmed for: %s", req.WalletAddress)

	// Check final approval status
	approved, err := handler.chain.IsApproved(req.WalletAddress)
	if err != nil {
		log.Printf("⚠️ Could not verify approval status: %v", err)
		approved = false // assume not approved if we can't check
	}

	render.JSON(w, r, map[string]interface{}{
		"status":   "confirmed",
		"message":  "User approval confirmed on blockchain",
		"tx_hash":  tx.Hash().Hex(),
		"approved": approved,
	})
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

func (handler *RequestHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := handler.db.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, users)
}

// AdminMiddleware ensures the authenticated user has the 'admin' role.
func (handler *RequestHandler) AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(auth.ClaimsKey).(*auth.Claims)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		user, err := handler.db.GetUserById(claims.UserID.String())
		if err != nil {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}

		if user.Role != models.RoleAdmin {
			http.Error(w, "Forbidden: Admin access required", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (handler *RequestHandler) UploadMetadata(w http.ResponseWriter, r *http.Request) {
	// 1. Parse Multipart Form (10 MB limit)
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	// 2. Retrieve the file from form data
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 3. Upload to IPFS via Pinata
	ipfsHash, err := ipfs.Upload(file, header.Filename)
	if err != nil {
		http.Error(w, "IPFS Upload Failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 4. Return the Hash to the Frontend
	render.JSON(w, r, map[string]string{
		"ipfs_hash": ipfsHash,
		"url":       "https://gateway.pinata.cloud/ipfs/" + ipfsHash,
	})
}
