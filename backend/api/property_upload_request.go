package api

import (
	"backend/auth"
	"backend/db/models"
	"backend/ipfs"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

// CreatePropertyUploadRequest handles POST /property-upload-requests
// Allows authenticated users to submit property upload requests
func (handler *RequestHandler) CreatePropertyUploadRequest(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if rec := recover(); rec != nil {
			log.Printf("❌ CreatePropertyUploadRequest: Panic recovered: %v", rec)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}()

	// Get authenticated user from context
	claims, ok := r.Context().Value(auth.ClaimsKey).(*auth.Claims)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Parse request
	payload, files, err := handler.parsePropertyUploadRequestRequest(r)
	if err != nil {
		log.Printf("❌ CreatePropertyUploadRequest: Request parsing failed: %v", err)
		http.Error(w, "Request Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("📋 CreatePropertyUploadRequest: Received request - User: %s, Name: %s, Symbol: %s, Valuation: %d, TokenSupply: %d, Files: %d",
		claims.UserID, payload.Name, payload.Symbol, payload.Valuation, payload.TokenSupply, len(files))

	// Process files and upload to IPFS
	dbDocs, mainHash, err := handler.processPropertyUploadRequestFiles(files)
	if err != nil {
		log.Printf("❌ CreatePropertyUploadRequest: File processing failed: %v", err)
		http.Error(w, "Upload Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("✅ CreatePropertyUploadRequest: Processed %d files, main hash: %s", len(dbDocs), mainHash)

	// Get user to get wallet address
	user, err := handler.db.GetUserById(claims.UserID.String())
	if err != nil {
		log.Printf("❌ CreatePropertyUploadRequest: User not found: %v", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Create property upload request
	request := models.PropertyUploadRequest{
		ID:            uuid.New(),
		UserID:        claims.UserID,
		WalletAddress: user.WalletAddress,
		Name:          payload.Name,
		Symbol:        payload.Symbol,
		Valuation:     float64(payload.Valuation),
		TokenSupply:   payload.TokenSupply,
		MetadataHash:  mainHash,
		Status:        models.ApprovalPending,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := handler.db.CreatePropertyUploadRequest(request); err != nil {
		log.Printf("❌ CreatePropertyUploadRequest: Database save failed: %v", err)
		http.Error(w, "Database Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("✅ CreatePropertyUploadRequest: Request saved to database - ID: %s", request.ID)

	// Link documents to request
	for i := range dbDocs {
		dbDocs[i].PropertyUploadRequestID = request.ID
		if err := handler.db.CreatePropertyUploadRequestDocument(dbDocs[i]); err != nil {
			log.Printf("⚠️ Failed to save document %d: %v", i, err)
		}
	}

	render.JSON(w, r, map[string]any{
		"status":  "success",
		"message": "Property upload request submitted successfully",
		"request_id": request.ID.String(),
		"files_count": len(dbDocs),
	})
}

// GetPropertyUploadRequests handles GET /property-upload-requests
// Returns all requests for admin, or user's own requests for regular users
func (handler *RequestHandler) GetPropertyUploadRequests(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(auth.ClaimsKey).(*auth.Claims)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get user to check role
	user, err := handler.db.GetUserById(claims.UserID.String())
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var requests []models.PropertyUploadRequest

	// Admin sees all requests, users see only their own
	if user.Role == models.RoleAdmin {
		requests, err = handler.db.GetPropertyUploadRequests()
		if err != nil {
			log.Printf("❌ GetPropertyUploadRequests: Failed to fetch requests: %v", err)
			http.Error(w, "Failed to fetch requests", http.StatusInternalServerError)
			return
		}
	} else {
		requests, err = handler.db.GetPropertyUploadRequestsByUser(user.WalletAddress)
		if err != nil {
			log.Printf("❌ GetPropertyUploadRequests: Failed to fetch user requests: %v", err)
			http.Error(w, "Failed to fetch requests", http.StatusInternalServerError)
			return
		}
	}

	render.JSON(w, r, requests)
}

// GetPropertyUploadRequest handles GET /property-upload-requests/{id}
// Returns a single request (admin can see any, users can see their own)
func (handler *RequestHandler) GetPropertyUploadRequest(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	claims, ok := r.Context().Value(auth.ClaimsKey).(*auth.Claims)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	request, err := handler.db.GetPropertyUploadRequestByID(id)
	if err != nil {
		http.Error(w, "Request not found", http.StatusNotFound)
		return
	}

	// Get user to check role
	user, err := handler.db.GetUserById(claims.UserID.String())
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Users can only see their own requests, admins can see any
	if user.Role != models.RoleAdmin && request.WalletAddress != user.WalletAddress {
		http.Error(w, "Forbidden: You can only view your own requests", http.StatusForbidden)
		return
	}

	render.JSON(w, r, request)
}

// ApprovePropertyUploadRequest handles POST /property-upload-requests/{id}/approve
// Admin-only: Creates the property on blockchain and updates request status
func (handler *RequestHandler) ApprovePropertyUploadRequest(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if rec := recover(); rec != nil {
			log.Printf("❌ ApprovePropertyUploadRequest: Panic recovered: %v", rec)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}()

	id := chi.URLParam(r, "id")

	// Get the request
	request, err := handler.db.GetPropertyUploadRequestByID(id)
	if err != nil {
		http.Error(w, "Request not found", http.StatusNotFound)
		return
	}

	// Check if already processed
	if request.Status != models.ApprovalPending {
		http.Error(w, fmt.Sprintf("Request already %s", request.Status), http.StatusBadRequest)
		return
	}

	log.Printf("📋 ApprovePropertyUploadRequest: Approving request %s - Name: %s, Owner: %s",
		id, request.Name, request.WalletAddress)

	// Check if chain service is available
	if handler.chain == nil {
		log.Printf("❌ ApprovePropertyUploadRequest: Blockchain service not available")
		http.Error(w, "Blockchain service not available", http.StatusServiceUnavailable)
		return
	}

	// Create property on blockchain using existing logic
	payload := PropertyPayload{
		OwnerAddress: request.WalletAddress,
		Name:         request.Name,
		Symbol:       request.Symbol,
		Valuation:    int64(request.Valuation),
		TokenSupply:  request.TokenSupply,
	}

	result, err := handler.createPropertyOnChain(&payload, request.MetadataHash)
	if err != nil {
		log.Printf("❌ ApprovePropertyUploadRequest: Blockchain transaction failed: %v", err)
		http.Error(w, "Blockchain Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("✅ ApprovePropertyUploadRequest: Blockchain transaction confirmed - Hash: %s", result.TxHash)

	// Create property record in database
	property := models.Property{
		ID:                  uuid.New(),
		Name:                result.PropertyName,
		OnchainAssetAddress: result.AssetAddress,
		OnchainTokenAddress: result.TokenAddress,
		OwnerWallet:         request.WalletAddress,
		MetadataHash:        request.MetadataHash,
		Valuation:           request.Valuation,
		Status:              models.StatusActive,
		TxHash:              result.TxHash,
		CreatedAt:           time.Now(),
	}

	if err := handler.db.CreateProperty(property); err != nil {
		log.Printf("❌ ApprovePropertyUploadRequest: Database save failed: %v", err)
		http.Error(w, "Database Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("✅ ApprovePropertyUploadRequest: Property saved to database - ID: %s", property.ID)

	// Copy documents from request to property
	// Get request documents (we need to fetch them, but for now we'll create them from request metadata)
	// In a real scenario, you'd fetch the request documents and copy them
	// For simplicity, we'll note that documents are already on IPFS via MetadataHash

	// Update request status to approved
	if err := handler.db.UpdatePropertyUploadRequestStatus(id, models.ApprovalApproved, ""); err != nil {
		log.Printf("⚠️ ApprovePropertyUploadRequest: Failed to update request status: %v", err)
		// Don't fail the request since property was created
	}

	render.JSON(w, r, map[string]any{
		"status":        "success",
		"message":       "Property created successfully and request approved",
		"request_id":    id,
		"property_id":   property.ID.String(),
		"tx_hash":       result.TxHash,
		"asset_address": result.AssetAddress,
		"token_address": result.TokenAddress,
	})
}

// RejectPropertyUploadRequest handles POST /property-upload-requests/{id}/reject
// Admin-only: Updates request status to rejected
func (handler *RequestHandler) RejectPropertyUploadRequest(w http.ResponseWriter, r *http.Request) {
	type RejectRequest struct {
		Reason string `json:"reason"`
	}

	var req RejectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}

	id := chi.URLParam(r, "id")

	// Get the request
	request, err := handler.db.GetPropertyUploadRequestByID(id)
	if err != nil {
		http.Error(w, "Request not found", http.StatusNotFound)
		return
	}

	// Check if already processed
	if request.Status != models.ApprovalPending {
		http.Error(w, fmt.Sprintf("Request already %s", request.Status), http.StatusBadRequest)
		return
	}

	log.Printf("📋 RejectPropertyUploadRequest: Rejecting request %s - Name: %s, Reason: %s",
		id, request.Name, req.Reason)

	// Update request status
	if err := handler.db.UpdatePropertyUploadRequestStatus(id, models.ApprovalRejected, req.Reason); err != nil {
		log.Printf("❌ RejectPropertyUploadRequest: Failed to update status: %v", err)
		http.Error(w, "Database Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]any{
		"status":     "success",
		"message":    "Property upload request rejected",
		"request_id": id,
	})
}

// ---------------------------------------------------------
// HELPER FUNCTIONS
// ---------------------------------------------------------

type PropertyUploadRequestPayload struct {
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Valuation   int64  `json:"valuation"`
	TokenSupply int64  `json:"token_supply"`
}

func (handler *RequestHandler) parsePropertyUploadRequestRequest(r *http.Request) (*PropertyUploadRequestPayload, []*multipart.FileHeader, error) {
	if err := r.ParseMultipartForm(20 << 20); err != nil {
		return nil, nil, fmt.Errorf("file too large or invalid format")
	}

	payloadData := r.FormValue("data")
	if payloadData == "" {
		return nil, nil, fmt.Errorf("missing 'data' field")
	}

	var payload PropertyUploadRequestPayload
	if err := json.Unmarshal([]byte(payloadData), &payload); err != nil {
		return nil, nil, fmt.Errorf("invalid JSON: %v", err)
	}

	if payload.Name == "" || payload.Valuation <= 0 {
		return nil, nil, fmt.Errorf("invalid payload fields")
	}

	files := r.MultipartForm.File["files"]
	if len(files) == 0 {
		return nil, nil, fmt.Errorf("at least one file is required")
	}

	return &payload, files, nil
}

func (handler *RequestHandler) processPropertyUploadRequestFiles(files []*multipart.FileHeader) ([]models.PropertyUploadRequestDocument, string, error) {
	var dbDocs []models.PropertyUploadRequestDocument
	var mainIpfsHash string

	for i, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, "", err
		}
		defer file.Close()

		docType := inferDocType(fileHeader.Filename)

		ext := filepath.Ext(fileHeader.Filename)
		rawName := strings.TrimSuffix(fileHeader.Filename, ext)
		uniqueName := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), rawName, ext)

		hash, err := ipfs.Upload(file, uniqueName)
		if err != nil {
			return nil, "", err
		}

		fullUrl := "https://gateway.pinata.cloud/ipfs/" + hash

		dbDocs = append(dbDocs, models.PropertyUploadRequestDocument{
			ID:         uuid.New(),
			FileUrl:    fullUrl,
			FileHash:   hash,
			Name:       uniqueName,
			Type:       docType,
			UploadedAt: time.Now(),
		})

		if i == 0 {
			mainIpfsHash = hash
		}
	}

	return dbDocs, mainIpfsHash, nil
}

