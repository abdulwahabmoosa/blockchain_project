package api

import (
	"backend/auth"
	"backend/blockchain"
	"backend/db"
	"backend/db/models"
	"backend/ipfs"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware" // CORS middleware
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10" // Ensure you ran: go get github.com/go-playground/validator/v10
)

// Initialize Validator
var validate = validator.New()

// Register custom validators
func init() {
	// eth_addr validator: checks if string starts with 0x and is 42 chars long (hex)
	validate.RegisterValidation("eth_addr", func(fl validator.FieldLevel) bool {
		addr := fl.Field().String()
		if len(addr) != 42 {
			return false
		}
		if !strings.HasPrefix(addr, "0x") {
			return false
		}
		// Check if remaining characters are valid hex
		for _, r := range addr[2:] {
			if !((r >= '0' && r <= '9') || (r >= 'a' && r <= 'f') || (r >= 'A' && r <= 'F')) {
				return false
			}
		}
		return true
	})
}

// --- Request Payloads with Validation Tags ---

type CreatePropertyRequest struct {
	OwnerAddress string  `json:"owner_address" validate:"required,eth_addr"`
	Name         string  `json:"name" validate:"required,min=3,max=100"`
	Symbol       string  `json:"symbol" validate:"required,alphanum,len=3-4"` // 3 or 4 chars, e.g. "PROP"
	DataHash     string  `json:"data_hash" validate:"required,printascii"`    // IPFS hash usually ascii
	Valuation    float64 `json:"valuation" validate:"required,gt=0"`          // Can be float from frontend
	TokenSupply  int64   `json:"token_supply" validate:"required,gt=0"`
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

	return &RequestHandler{db, chain}
}

func (handler *RequestHandler) Start() {
	r := chi.NewRouter()

	// CORS middleware configuration
	// r.Use(cors.Handler(cors.Options{
	// 	AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:3000"}, // Allow frontend dev server and backend
	// 	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	// 	ExposedHeaders:   []string{"Link"},
	// 	AllowCredentials: true,
	// 	MaxAge:           300, // Maximum value not ignored by any of major browsers
	// }))

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // Recover from panics and return 500
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Post("/login", handler.Login)
	r.Post("/register", handler.RegisterUser)

	// Temporarily move upload outside auth for testing
	r.Post("/upload", handler.UploadMetadata)

	r.Group(func(r chi.Router) {
		r.Use(auth.Middleware)

		// Property Routes
		r.Get("/properties", handler.GetProperties)
		r.Get("/properties/{id}", handler.GetProperty)
		r.Get("/properties/{id}/metadata", handler.GetPropertyMetadata)
		r.Get("/properties/{id}/token-balance/{wallet}", handler.GetPropertyTokenBalance)
		r.Post("/properties/{id}/transfer", handler.TransferPropertyTokens)

		r.Put("/users/me", handler.UpdateUserInfo)
		r.Delete("/users/me", handler.DeleteUser)
		r.Post("/users/me/reset-password", handler.ResetPassword)

		// Admin Routes
		r.Group(func(r chi.Router) {
			r.Use(handler.AdminMiddleware)
			r.Get("/users", handler.GetUsers)
			r.Post("/approve-user", handler.ApproveUser)
			r.Post("/users/approval", handler.UpdateUserApproval)
			r.Post("/properties", handler.CreateProperty)
			r.Post("/properties/approval", handler.UpdatePropertyApproval)
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

	handler.gracefulShutdown(srv)

}

func (handler *RequestHandler) gracefulShutdown(srv *http.Server) {
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

// Login is in auth.go

func (handler *RequestHandler) UpdateUserInfo(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(auth.ClaimsKey).(*auth.Claims)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	type UpdateRequest struct {
		Name  string `json:"name"`
		Email string `json:"email" validate:"omitempty,email"`
	}

	var req UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}

	if req.Email != "" {
		if err := validate.Var(req.Email, "email"); err != nil {
			http.Error(w, "Invalid Email Format", http.StatusBadRequest)
			return
		}
	}

	if err := handler.db.UpdateUserInfo(claims.UserID.String(), req.Name, req.Email); err != nil {
		http.Error(w, "DB Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]string{
		"status":  "success",
		"message": "Profile updated successfully",
	})
}

func (handler *RequestHandler) UpdateUserApproval(w http.ResponseWriter, r *http.Request) {
	type UserApprovalRequest struct {
		WalletAddress string                `json:"wallet_address"`
		Status        models.ApprovalStatus `json:"status"`
	}

	var req UserApprovalRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}

	if req.WalletAddress == "" {
		http.Error(w, "wallet_address is required", http.StatusBadRequest)
		return
	}

	if req.Status != models.ApprovalApproved && req.Status != models.ApprovalRejected {
		http.Error(w, "status must be 'approved' or 'rejected'", http.StatusBadRequest)
		return
	}

	tx, err := handler.chain.ApproveUser(req.WalletAddress)
	if err != nil {
		http.Error(w, "Blockchain error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 2. WAIT for mining
	log.Printf("⏳ Waiting for approval tx to mine: %s", tx.Hash().Hex())
	receipt, err := bind.WaitMined(r.Context(), handler.chain.Client, tx)
	if err != nil || receipt.Status == 0 {
		http.Error(w, "Transaction failed on-chain", http.StatusInternalServerError)
		return
	}

	// 3. Update DB immediately
	err = handler.db.UpdateUserApproval(req.WalletAddress, "approved")
	if err != nil {
		log.Printf("❌ Failed to update DB after mining: %v", err)
	}

	render.JSON(w, r, map[string]string{"tx_hash": tx.Hash().Hex(), "": "confirmed"})
}

// DeleteUser handles DELETE /users/me
func (handler *RequestHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(auth.ClaimsKey).(*auth.Claims)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := handler.db.DeleteUser(claims.UserID.String()); err != nil {
		http.Error(w, "DB Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]string{
		"status":  "success",
		"message": "Account deleted successfully",
	})
}

// ResetPassword handles POST /users/me/reset-password
func (handler *RequestHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(auth.ClaimsKey).(*auth.Claims)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	type ResetRequest struct {
		OldPassword string `json:"old_password" validate:"required"`
		NewPassword string `json:"new_password" validate:"required,min=6"`
	}

	var req ResetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}

	user, err := handler.db.GetUserById(claims.UserID.String())
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if !auth.CheckPasswordHash(req.OldPassword, user.PasswordHash) {
		http.Error(w, "Incorrect old password", http.StatusUnauthorized)
		return
	}

	// 3. Hash New Password
	newHash, err := auth.HashPassword(req.NewPassword)
	if err != nil {
		http.Error(w, "Hashing Error", http.StatusInternalServerError)
		return
	}

	// 4. Update DB
	if err := handler.db.UpdatePassword(claims.UserID.String(), newHash); err != nil {
		http.Error(w, "DB Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]string{
		"status":  "success",
		"message": "Password updated successfully",
	})
}

// RegisterUser is in auth.go

// CreateProperty is now in property.go - uses multipart form data with files

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

	// // Check if chain service is available
	// if handler.chain == nil {
	// 	log.Printf("❌ Chain service not initialized - blockchain functionality disabled")
	// 	http.Error(w, "Blockchain service not available - check environment configuration", http.StatusServiceUnavailable)
	// 	return
	// }
	// chain, err := blockchain.NewChainServiceEnv()
	// if err != nil {
	// 	log.Printf("❌ Chain init failed: %v", err)
	// } else {
	// 	log.Printf("✅ Chain service initialized successfully")
	// }
	// handler.chain = chain

	// Check if Approval contract is deployed
	if handler.chain.Approval == nil {
		log.Printf("❌ Approval contract not deployed - cannot approve users")
		http.Error(w, "Blockchain Error: approval contract not deployed - deploy contracts to enable blockchain functionality", http.StatusServiceUnavailable)
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

// GetProperties and GetProperty are now in property.go

func (handler *RequestHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := handler.db.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, users)
}

// UpdatePropertyApproval is now in property.go with blockchain integration

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

func (handler *RequestHandler) GetPropertyMetadata(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	prop, err := handler.db.GetPropertyByID(id)
	if err != nil {
		http.Error(w, "Property not found", http.StatusNotFound)
		return
	}

	if prop.MetadataHash == "" {
		http.Error(w, "Property metadata not available", http.StatusNotFound)
		return
	}

	// Fetch metadata from IPFS
	metadataURL := fmt.Sprintf("https://gateway.pinata.cloud/ipfs/%s", prop.MetadataHash)
	resp, err := http.Get(metadataURL)
	if err != nil {
		log.Printf("Failed to fetch metadata from IPFS: %v", err)
		http.Error(w, "Failed to fetch metadata from IPFS", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("IPFS gateway returned status %d", resp.StatusCode)
		http.Error(w, "Metadata not found on IPFS", http.StatusNotFound)
		return
	}

	var metadata map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&metadata); err != nil {
		log.Printf("Failed to parse metadata JSON: %v", err)
		http.Error(w, "Invalid metadata format", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, metadata)
}

func (handler *RequestHandler) GetPropertyTokenBalance(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	wallet := chi.URLParam(r, "wallet")

	prop, err := handler.db.GetPropertyByID(id)
	if err != nil {
		http.Error(w, "Property not found", http.StatusNotFound)
		return
	}

	if handler.chain == nil {
		http.Error(w, "Blockchain service not available", http.StatusServiceUnavailable)
		return
	}

	balance, err := handler.chain.GetTokenBalance(prop.OnchainTokenAddress, wallet)
	if err != nil {
		log.Printf("Failed to get token balance: %v", err)
		http.Error(w, "Failed to get token balance: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert from wei (assuming 18 decimals)
	balanceFloat := new(big.Float).SetInt(balance)
	divisor := new(big.Float).SetInt(big.NewInt(1000000000000000000)) // 10^18
	balanceFloat.Quo(balanceFloat, divisor)

	render.JSON(w, r, map[string]string{
		"balance": balanceFloat.Text('f', 18),
	})
}

func (handler *RequestHandler) TransferPropertyTokens(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	type TransferRequest struct {
		ToAddress string `json:"to_address" validate:"required,eth_addr"`
		Amount    string `json:"amount" validate:"required"`
	}

	var req TransferRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(req); err != nil {
		http.Error(w, "Validation Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	prop, err := handler.db.GetPropertyByID(id)
	if err != nil {
		http.Error(w, "Property not found", http.StatusNotFound)
		return
	}

	if handler.chain == nil {
		http.Error(w, "Blockchain service not available", http.StatusServiceUnavailable)
		return
	}

	// Parse amount (assuming it's in token units, need to convert to wei)
	amountFloat, _, err := big.ParseFloat(req.Amount, 10, 256, big.ToNearestEven)
	if err != nil {
		http.Error(w, "Invalid amount format", http.StatusBadRequest)
		return
	}

	weiMultiplier := big.NewFloat(1000000000000000000) // 10^18
	amountFloat.Mul(amountFloat, weiMultiplier)
	amountBig, _ := amountFloat.Int(nil)

	tx, err := handler.chain.TransferTokens(prop.OnchainTokenAddress, req.ToAddress, amountBig)
	if err != nil {
		log.Printf("Failed to transfer tokens: %v", err)
		http.Error(w, "Transfer failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]string{
		"status":  "pending",
		"tx_hash": tx.Hash().Hex(),
		"message": "Transfer transaction submitted",
	})
}
