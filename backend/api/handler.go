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
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10" // Ensure you ran: go get github.com/go-playground/validator/v10
	"github.com/google/uuid"
)

// Initialize Validator
var validate = validator.New()

// Register custom validaters
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
// structs for API request validation

type CreatePropertyRequest struct {
	OwnerAddress string  `json:"owner_address" validate:"required,eth_addr"`
	Name         string  `json:"name" validate:"required,min=3,max=100"`
	Symbol       string  `json:"symbol" validate:"required,alphanum,len=3-4"` // 3 or 4 chars, e.g. "PROP"
	DataHash     string  `json:"data_hash" validate:"required,printascii"`    // IPFS hash usually ascii
	Valuation    float64 `json:"valuation" validate:"required,gt=0"`          // Can be float from frontend
	TokenSupply  int64   `json:"token_supply" validate:"required,gt=0"` // Token supply amount as integer
}

type DistributeRevenueRequest struct {
	TokenAddress      string `json:"token_address" validate:"required,eth_addr"`
	StablecoinAddress string `json:"stablecoin_address" validate:"required,eth_addr"`
	Amount            int64  `json:"amount" validate:"required,gt=0"`
}

type ApproveUserRequest struct {
	WalletAddress string `json:"wallet_address" validate:"required,eth_addr"`
}

// RequestHandler - main API handler struct
// contains database and blockchain service references
type RequestHandler struct {
	db    *db.Database      // database connection
	chain *blockchain.ChainService // blockchain service (optional)
}

// NewRequestHandler - create new API handler instance
func NewRequestHandler(db *db.Database, chain *blockchain.ChainService) *RequestHandler {
	// Setup blockchain event listeners (optional - won't crash if subscriptions fail)

	return &RequestHandler{db, chain}
}

// Start - setup routes and start HTTP server
func (handler *RequestHandler) Start() {
	r := chi.NewRouter()

	// CORS middleware configuration
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*.vercel.app", "http://localhost:5173", "http://localhost:3000", "http://127.0.0.1:5173"}, // Allow Vercel domains and local development
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Logger)
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					log.Printf("PANIC RECOVERED in middleware: %v, Path: %s, Method: %s", rec, r.URL.Path, r.Method)
					http.Error(w, "Internal server error", http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	})
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
		r.Get("/properties/{id}/token-stats", handler.GetPropertyTokenStats)
		r.Post("/properties/{id}/transfer", handler.TransferPropertyTokens)
		r.Post("/properties/{id}/purchase", handler.CreateTokenPurchase)
		r.Get("/properties/{id}/pending-purchases", handler.GetPendingTokenPurchases)
		r.Post("/properties/{id}/purchases/{purchaseId}/approve", handler.ApproveTokenPurchase)
		r.Post("/properties/{id}/purchases/{purchaseId}/update-tx", handler.UpdateTokenPurchaseTxHash)
		
		// User routes - use Route() to ensure proper sub-path matching
		r.Route("/users/me", func(r chi.Router) {
			r.Get("/purchases", handler.GetMyTokenPurchases)
			r.Get("/pending-transfers", handler.GetMyPendingTransfers)
			r.Post("/reset-password", handler.ResetPassword)
			r.Get("/", handler.GetCurrentUser) // "/" matches /users/me exactly
			r.Put("/", handler.UpdateUserInfo)
			r.Delete("/", handler.DeleteUser)
		})

		// Property Upload Request Routes (authenticated users)
		r.Post("/property-upload-requests", handler.CreatePropertyUploadRequest)
		r.Get("/property-upload-requests", handler.GetPropertyUploadRequests)
		r.Get("/property-upload-requests/user", handler.GetPropertyUploadRequests)
		r.Get("/property-upload-requests/{id}", handler.GetPropertyUploadRequest)

		// Admin Routes
		r.Group(func(r chi.Router) {
			r.Use(handler.AdminMiddleware)
			r.Get("/users", handler.GetUsers)
			r.Post("/approve-user", handler.ApproveUser)
			r.Post("/users/approval", handler.UpdateUserApproval)
			r.Post("/properties", handler.CreateProperty)
			r.Post("/properties/approval", handler.UpdatePropertyApproval)
			r.Post("/revenue/distribute", handler.DistributeRevenue)
			r.Post("/property-upload-requests/{id}/approve", handler.ApprovePropertyUploadRequest)
			r.Post("/property-upload-requests/{id}/reject", handler.RejectPropertyUploadRequest)
		})
	})

	srv := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	// Start Server in a Goroutine so it dont block
	go func() {
		log.Println("Starting server on port 3000...")
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
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exited successfully")
}

// Login is in auth.go

// GetCurrentUser - get authenticated user profile
// GET /users/me - requires authentication
func (handler *RequestHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(auth.ClaimsKey).(*auth.Claims)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := handler.db.GetUserById(claims.UserID.String())
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	render.JSON(w, r, user)
}

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
	log.Printf("Waiting for approval tx to mine: %s", tx.Hash().Hex())
	receipt, err := bind.WaitMined(r.Context(), handler.chain.Client, tx)
	if err != nil || receipt.Status == 0 {
		http.Error(w, "Transaction failed on-chain", http.StatusInternalServerError)
		return
	}

	// 3. Update DB immediately
	err = handler.db.UpdateUserApproval(req.WalletAddress, "approved")
	if err != nil {
		log.Printf("Failed to update DB after mining: %v", err)
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

// DistributeRevenue handles revenue distribution to property tokens

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

// ApproveUser - approve user for platform access
// POST /approve-user - admin only, requires wallet_address
func (handler *RequestHandler) ApproveUser(w http.ResponseWriter, r *http.Request) {
	// Middleware has already checked Admin role

	var req ApproveUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Failed to decode request: %v", err)
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	// VALIDATION
	if err := validate.Struct(req); err != nil {
		log.Printf("Validation failed: %v", err)
		http.Error(w, "Validation Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	// // Check if chain service is available
	// if handler.chain == nil {
	// 	log.Printf("Chain service not initialized - blockchain functionality disabled")
	// 	http.Error(w, "Blockchain service not available - check environment configuration", http.StatusServiceUnavailable)
	// 	return
	// }
	// chain, err := blockchain.NewChainServiceEnv()
	// if err != nil {
	// 	log.Printf("Chain init failed: %v", err)
	// } else {
	// 	log.Printf("Chain service initialized successfully")
	// }
	// handler.chain = chain

	// Check if Approval contract is deployed
	if handler.chain.Approval == nil {
		log.Printf("Approval contract not deployed - cannot approve users")
		http.Error(w, "Blockchain Error: approval contract not deployed - deploy contracts to enable blockchain functionality", http.StatusServiceUnavailable)
		return
	}

	log.Printf("Approving user on blockchain: %s", req.WalletAddress)
	tx, err := handler.chain.ApproveUser(req.WalletAddress)
	if err != nil {
		log.Printf("Blockchain approval failed for %s: %v", req.WalletAddress, err)
		http.Error(w, "Blockchain Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Blockchain approval confirmed for: %s", req.WalletAddress)

	// Check final approval status
	approved, err := handler.chain.IsApproved(req.WalletAddress)
	if err != nil {
		log.Printf("Warning: Could not verify approval status: %v", err)
		approved = false // assume not approved if we can't check
	}

	render.JSON(w, r, map[string]interface{}{
		"status":   "confirmed",
		"message":  "User approval confirmed on blockchain",
		"tx_hash":  tx.Hash().Hex(),
		"approved": approved,
	})
}

// GetUsers - get all users in system
// GET /users - admin only endpoint
func (handler *RequestHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := handler.db.GetAllUsers()
	if err != nil {
		log.Printf("GetUsers error: %v", err)
		http.Error(w, "Failed to fetch users: "+err.Error(), http.StatusInternalServerError)
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

// UploadMetadata - upload file to IPFS
// POST /upload - accepts multipart form with file
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

func (handler *RequestHandler) GetPropertyTokenStats(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	prop, err := handler.db.GetPropertyByID(id)
	if err != nil {
		http.Error(w, "Property not found", http.StatusNotFound)
		return
	}

	if handler.chain == nil {
		http.Error(w, "Blockchain service not available", http.StatusServiceUnavailable)
		return
	}

	// Get total supply from blockchain
	totalSupply, err := handler.chain.GetTotalSupply(prop.OnchainTokenAddress)
	if err != nil {
		log.Printf("Failed to get total supply: %v", err)
		http.Error(w, "Failed to get total supply: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert total supply from wei to token units
	totalSupplyFloat := new(big.Float).SetInt(totalSupply)
	divisor := new(big.Float).SetInt(big.NewInt(1000000000000000000)) // 10^18
	totalSupplyFloat.Quo(totalSupplyFloat, divisor)
	total, _ := totalSupplyFloat.Float64()

	// Get total sold from database
	totalSold, err := handler.db.GetTokenStatsByProperty(id)
	if err != nil {
		log.Printf("Failed to get token stats: %v", err)
		// If database query fails, default to 0 sold
		totalSold = 0
	}

	// Calculate available
	available := total - totalSold
	if available < 0 {
		available = 0 // Ensure non-negative
	}

	// Calculate percentage sold
	var percentageSold float64
	if total > 0 {
		percentageSold = (totalSold / total) * 100
	}

	render.JSON(w, r, map[string]interface{}{
		"total":           total,
		"sold":            totalSold,
		"available":       available,
		"percentage_sold": percentageSold,
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

	txHash := tx.Hash().Hex()

	// Record the purchase in database (non-blocking - if it fails, log but don't fail the request)
	// Convert amountBig back to token units (divide by 10^18) for storage
	amountFloatForDB := new(big.Float).SetInt(amountBig)
	divisor := new(big.Float).SetInt(big.NewInt(1000000000000000000)) // 10^18
	amountFloatForDB.Quo(amountFloatForDB, divisor)
	amountStr := amountFloatForDB.Text('f', 18) // Format with 18 decimal places

	propertyUID, _ := uuid.Parse(id)
	purchase := models.TokenPurchase{
		ID:          uuid.New(),
		PropertyID:  propertyUID,
		BuyerWallet: req.ToAddress,
		Amount:      amountStr,
		TokenTxHash: txHash,
		CreatedAt:   time.Now(),
	}

	if err := handler.db.CreateTokenPurchase(purchase); err != nil {
		// Log error but don't fail the request - tokens are already transferred
		log.Printf("Warning: Failed to record token purchase in database: %v (Transaction: %s)", err, txHash)
	} else {
		log.Printf("Token purchase recorded: Property=%s, Buyer=%s, Amount=%s, TX=%s", id, req.ToAddress, amountStr, txHash)
	}

	render.JSON(w, r, map[string]string{
		"status":  "pending",
		"tx_hash": txHash,
		"message": "Transfer transaction submitted",
	})
}

// CreateTokenPurchase handles POST /properties/{id}/purchase
// Records a token purchase after buyer sends payment (before owner approves transfer)
func (handler *RequestHandler) CreateTokenPurchase(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	type PurchaseRequest struct {
		BuyerWallet   string `json:"buyer_wallet" validate:"required,eth_addr"`
		Amount        string `json:"amount" validate:"required"` // Token amount in token units
		PaymentTxHash string `json:"payment_tx_hash" validate:"required"` // ETH payment transaction hash
		PurchasePrice string `json:"purchase_price" validate:"required"` // ETH paid
	}

	var req PurchaseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(req); err != nil {
		http.Error(w, "Validation Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Verify property exists
	_, err := handler.db.GetPropertyByID(id)
	if err != nil {
		http.Error(w, "Property not found", http.StatusNotFound)
		return
	}

	propertyUID, _ := uuid.Parse(id)
	purchase := models.TokenPurchase{
		ID:            uuid.New(),
		PropertyID:    propertyUID,
		BuyerWallet:   req.BuyerWallet,
		Amount:        req.Amount,
		PaymentTxHash: req.PaymentTxHash,
		TokenTxHash:   "pending", // Will be updated when owner approves
		PurchasePrice: req.PurchasePrice,
		CreatedAt:     time.Now(),
	}

	if err := handler.db.CreateTokenPurchase(purchase); err != nil {
		log.Printf("Failed to create token purchase: %v", err)
		http.Error(w, "Database Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Token purchase recorded: Property=%s, Buyer=%s, Amount=%s, PaymentTX=%s", id, req.BuyerWallet, req.Amount, req.PaymentTxHash)

	render.JSON(w, r, map[string]interface{}{
		"status":      "success",
		"purchase_id": purchase.ID.String(),
		"message":     "Purchase recorded. Waiting for owner approval.",
	})
}

// GetPendingTokenPurchases handles GET /properties/{id}/pending-purchases
// Returns all pending token purchases for a property (for owner to approve)
func (handler *RequestHandler) GetPendingTokenPurchases(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	// Get property to verify it exists
	prop, err := handler.db.GetPropertyByID(id)
	if err != nil {
		http.Error(w, "Property not found", http.StatusNotFound)
		return
	}

	// Get authenticated user
	claims, ok := r.Context().Value(auth.ClaimsKey).(*auth.Claims)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get user to check if they're the owner
	user, err := handler.db.GetUserById(claims.UserID.String())
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Verify user is the property owner
	if user.WalletAddress == "" || user.WalletAddress != prop.OwnerWallet {
		http.Error(w, "Forbidden: Only the property owner can view pending purchases", http.StatusForbidden)
		return
	}

	// Get pending purchases
	purchases, err := handler.db.GetPendingTokenPurchasesByProperty(id)
	if err != nil {
		log.Printf("Failed to get pending purchases: %v", err)
		http.Error(w, "Database Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, purchases)
}

// ApproveTokenPurchase handles POST /properties/{id}/purchases/{purchaseId}/approve
// Owner approves purchase and transfers tokens, then updates the purchase record
func (handler *RequestHandler) ApproveTokenPurchase(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	purchaseId := chi.URLParam(r, "purchaseId")

	// Get property
	prop, err := handler.db.GetPropertyByID(id)
	if err != nil {
		http.Error(w, "Property not found", http.StatusNotFound)
		return
	}

	// Get authenticated user
	claims, ok := r.Context().Value(auth.ClaimsKey).(*auth.Claims)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get user to check if they're the owner
	user, err := handler.db.GetUserById(claims.UserID.String())
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Verify user is the property owner
	if user.WalletAddress == "" || user.WalletAddress != prop.OwnerWallet {
		http.Error(w, "Forbidden: Only the property owner can approve purchases", http.StatusForbidden)
		return
	}

	// Get the purchase record
	purchases, err := handler.db.GetPendingTokenPurchasesByProperty(id)
	if err != nil {
		http.Error(w, "Database Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var purchase *models.TokenPurchase
	for i := range purchases {
		if purchases[i].ID.String() == purchaseId {
			purchase = &purchases[i]
			break
		}
	}

	if purchase == nil {
		http.Error(w, "Purchase not found or already approved", http.StatusNotFound)
		return
	}

	if handler.chain == nil {
		http.Error(w, "Blockchain service not available", http.StatusServiceUnavailable)
		return
	}

	// Parse amount (convert from token units to wei)
	amountFloat, _, err := big.ParseFloat(purchase.Amount, 10, 256, big.ToNearestEven)
	if err != nil {
		http.Error(w, "Invalid amount format", http.StatusBadRequest)
		return
	}

	weiMultiplier := big.NewFloat(1000000000000000000) // 10^18
	amountFloat.Mul(amountFloat, weiMultiplier)
	amountBig, _ := amountFloat.Int(nil)

	// Transfer tokens from owner to buyer
	// Note: This requires the owner's wallet to have tokens and be connected
	// The frontend will handle the actual transfer using the owner's signer
	// This endpoint just updates the database after the frontend confirms the transfer

	// For now, we'll return the purchase details and let the frontend handle the transfer
	// Then the frontend will call back to update the token_tx_hash
	render.JSON(w, r, map[string]interface{}{
		"status":      "ready",
		"purchase":    purchase,
		"message":     "Purchase ready for approval. Frontend should transfer tokens and then update the record.",
		"token_address": prop.OnchainTokenAddress,
		"buyer_address": purchase.BuyerWallet,
		"amount_wei":    amountBig.String(),
	})
}

// UpdateTokenPurchaseTxHash handles POST /properties/{id}/purchases/{purchaseId}/update-tx
// Updates the token transfer transaction hash after owner approves and transfers tokens
func (handler *RequestHandler) UpdateTokenPurchaseTxHash(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	purchaseId := chi.URLParam(r, "purchaseId")

	type UpdateRequest struct {
		TokenTxHash string `json:"token_tx_hash" validate:"required"`
	}

	var req UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(req); err != nil {
		http.Error(w, "Validation Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Get property
	prop, err := handler.db.GetPropertyByID(id)
	if err != nil {
		http.Error(w, "Property not found", http.StatusNotFound)
		return
	}

	// Get authenticated user
	claims, ok := r.Context().Value(auth.ClaimsKey).(*auth.Claims)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get user to check if they're the owner
	user, err := handler.db.GetUserById(claims.UserID.String())
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Verify user is the property owner
	if user.WalletAddress == "" || user.WalletAddress != prop.OwnerWallet {
		http.Error(w, "Forbidden: Only the property owner can update purchases", http.StatusForbidden)
		return
	}

	// Update the purchase record
	if err := handler.db.UpdateTokenPurchaseTxHash(purchaseId, req.TokenTxHash); err != nil {
		log.Printf("Failed to update token purchase: %v", err)
		http.Error(w, "Database Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Token purchase updated: PurchaseID=%s, TokenTX=%s", purchaseId, req.TokenTxHash)

	render.JSON(w, r, map[string]interface{}{
		"status":  "success",
		"message": "Purchase record updated successfully",
	})
}

// GetMyTokenPurchases handles GET /users/me/purchases
// Returns all token purchases made by the authenticated user (as buyer)
func (handler *RequestHandler) GetMyTokenPurchases(w http.ResponseWriter, r *http.Request) {
	log.Printf("Debug GetMyTokenPurchases: Handler called - Method: %s, Path: %s", r.Method, r.URL.Path)
	
	// Get authenticated user
	claims, ok := r.Context().Value(auth.ClaimsKey).(*auth.Claims)
	if !ok {
		log.Printf("GetMyTokenPurchases: Unauthorized - no claims in context")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	log.Printf("GetMyTokenPurchases: User ID: %s", claims.UserID.String())

	// Get user to get wallet address
	user, err := handler.db.GetUserById(claims.UserID.String())
	if err != nil {
		log.Printf("GetMyTokenPurchases: User not found: %v", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if user.WalletAddress == "" {
		log.Printf("GetMyTokenPurchases: User wallet address not found")
		http.Error(w, "User wallet address not found", http.StatusBadRequest)
		return
	}

	log.Printf("Debug GetMyTokenPurchases: Fetching purchases for wallet: %s", user.WalletAddress)
	
	// Get all purchases by this buyer
	purchases, err := handler.db.GetTokenPurchasesByBuyer(user.WalletAddress)
	if err != nil {
		log.Printf("Failed to get token purchases: %v", err)
		http.Error(w, "Database Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Always return an array, even if empty
	if purchases == nil {
		purchases = []models.TokenPurchase{}
	}

	log.Printf("Debug GetMyTokenPurchases: Returning %d purchases", len(purchases))
	render.JSON(w, r, purchases)
}

// GetMyPendingTransfers handles GET /users/me/pending-transfers
// Returns all pending token transfers for properties owned by the authenticated user
func (handler *RequestHandler) GetMyPendingTransfers(w http.ResponseWriter, r *http.Request) {
	log.Printf("GetMyPendingTransfers: Handler called - Method: %s, Path: %s", r.Method, r.URL.Path)

	// Get authenticated user
	claims, ok := r.Context().Value(auth.ClaimsKey).(*auth.Claims)
	if !ok {
		log.Printf("GetMyPendingTransfers: Unauthorized - no claims in context")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	log.Printf("GetMyPendingTransfers: User ID: %s", claims.UserID.String())

	// Get user to get wallet address
	user, err := handler.db.GetUserById(claims.UserID.String())
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if user.WalletAddress == "" {
		http.Error(w, "User wallet address not found", http.StatusBadRequest)
		return
	}

	// Get all properties owned by this user
	allProperties, err := handler.db.GetAllProperties()
	if err != nil {
		log.Printf("Failed to get properties: %v", err)
		http.Error(w, "Database Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Filter properties owned by user
	userWalletLower := strings.ToLower(user.WalletAddress)
	var ownedProperties []models.Property
	for _, prop := range allProperties {
		if prop.OwnerWallet != "" && strings.ToLower(prop.OwnerWallet) == userWalletLower {
			ownedProperties = append(ownedProperties, prop)
		}
	}

	log.Printf("GetMyPendingTransfers: Found %d owned properties", len(ownedProperties))

	// Get pending purchases for all owned properties
	var allPendingTransfers []models.TokenPurchase
	for _, prop := range ownedProperties {
		pending, err := handler.db.GetPendingTokenPurchasesByProperty(prop.ID.String())
		if err != nil {
			log.Printf("Warning: Failed to get pending purchases for property %s: %v", prop.ID, err)
			continue
		}
		log.Printf("GetMyPendingTransfers: Property %s has %d pending purchases", prop.ID, len(pending))
		allPendingTransfers = append(allPendingTransfers, pending...)
	}

	log.Printf("GetMyPendingTransfers: Returning %d total pending transfers", len(allPendingTransfers))
	
	// Always return an array, even if empty
	if allPendingTransfers == nil {
		allPendingTransfers = []models.TokenPurchase{}
	}
	
	render.JSON(w, r, allPendingTransfers)
}
