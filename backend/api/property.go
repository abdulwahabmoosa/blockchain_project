package api

import (
	"backend/blockchain"
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

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

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

func (handler *RequestHandler) CreateProperty(w http.ResponseWriter, r *http.Request) {
	// Ensure we always send a response, even on panik
	defer func() {
		if rec := recover(); rec != nil {
			log.Printf("CreateProperty: Panic recovered: %v", rec)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}()

	payload, files, err := handler.parseCreatePropertyRequest(r)
	if err != nil {
		log.Printf("CreateProperty: Request parsing failed: %v", err)
		http.Error(w, "Request Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("CreateProperty: Received request - Owner: %s, Name: %s, Symbol: %s, Valuation: %d, TokenSupply: %d, Files: %d",
		payload.OwnerAddress, payload.Name, payload.Symbol, payload.Valuation, payload.TokenSupply, len(files))

	// Check if chain service is available
	if handler.chain == nil {
		log.Printf("CreateProperty: Blockchain service not available")
		http.Error(w, "Blockchain service not available - check environment configuration", http.StatusServiceUnavailable)
		return
	}

	dbDocs, mainHash, err := handler.processPropertyFiles(files)
	if err != nil {
		log.Printf("CreateProperty: File processing failed: %v", err)
		http.Error(w, "Upload Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("CreateProperty: Processed %d files, main hash: %s", len(dbDocs), mainHash)

	// Submit to blockchain and wait for confirmation
	result, err := handler.createPropertyOnChain(payload, mainHash)
	if err != nil {
		log.Printf("CreateProperty: Blockchain transaction failed: %v", err)
		http.Error(w, "Blockchain Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("CreateProperty: Blockchain transaction confirmed - Hash: %s", result.TxHash)
	log.Printf("CreateProperty: Asset: %s, Token: %s", result.AssetAddress, result.TokenAddress)

	// Create property record in databse
	property := models.Property{
		ID:                  uuid.New(),
		Name:                result.PropertyName,
		OnchainAssetAddress: result.AssetAddress,
		OnchainTokenAddress: result.TokenAddress,
		OwnerWallet:         payload.OwnerAddress,
		MetadataHash:        mainHash,
		Valuation:           float64(payload.Valuation),
		Status:              models.StatusActive,
		TxHash:              result.TxHash,
		CreatedAt:           time.Now(),
	}

	if err := handler.db.CreateProperty(property); err != nil {
		log.Printf("CreateProperty: Database save failed: %v", err)
		// Property exists on blockchain but not in DB - log error but don't fail the request
		// The event listener can pick it up later
		log.Printf("Warning: Property created on blockchain but DB save failed. Event listener will retry.")
	} else {
		log.Printf("CreateProperty: Property saved to database - ID: %s", property.ID)
	}

	// Link documents to property
	for i := range dbDocs {
		dbDocs[i].PropertyID = property.ID
		if err := handler.db.CreatePropertyDocument(dbDocs[i]); err != nil {
			log.Printf("Warning: Failed to save document %d: %v", i, err)
		}
	}

	render.JSON(w, r, map[string]any{
		"status":        "success",
		"tx_hash":       result.TxHash,
		"property_id":   property.ID.String(),
		"asset_address": result.AssetAddress,
		"token_address": result.TokenAddress,
		"files_count":   len(dbDocs),
		"message":       "Property created successfully on blockchain and saved to database.",
	})
}

func (handler *RequestHandler) UpdatePropertyApproval(w http.ResponseWriter, r *http.Request) {
	type ReviewRequest struct {
		PropertyID string                `json:"property_id"`
		Status     models.ApprovalStatus `json:"status"`
	}

	var req ReviewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}

	if req.PropertyID == "" {
		http.Error(w, "property_id is required", http.StatusBadRequest)
		return
	}

	if req.Status != models.ApprovalApproved && req.Status != models.ApprovalRejected {
		http.Error(w, "status must be 'approved' or 'rejected'", http.StatusBadRequest)
		return
	}

	// Get property to get the asset address for blockchain calls
	prop, err := handler.db.GetPropertyByID(req.PropertyID)
	if err != nil {
		http.Error(w, "Property not found", http.StatusNotFound)
		return
	}

	// Call blockchain if available
	var tx *types.Transaction
	if handler.chain != nil {
		switch req.Status {
		case models.ApprovalApproved:
			tx, err = handler.chain.ApproveProperty(prop.OnchainAssetAddress)
			if err != nil {
				log.Printf("Warning: Blockchain approval failed: %v", err)
				// Continue with DB update even if blockchain fails
			} else {
				log.Printf("Blockchain approval transaction: %s", tx.Hash().Hex())
			}
		case models.ApprovalRejected:
			tx, err = handler.chain.RejectProperty(prop.OnchainAssetAddress)
			if err != nil {
				log.Printf("Warning: Blockchain rejection failed: %v", err)
				// Continue with DB update even if blockchain fails
			} else {
				log.Printf("Blockchain rejection transaction: %s", tx.Hash().Hex())
			}
		}
	}

	// Update database
	if err := handler.db.UpdatePropertyApproval(req.PropertyID, req.Status); err != nil {
		if err.Error() == "property_id is required" {
			http.Error(w, "Property not found", http.StatusNotFound)
			return
		}
		http.Error(w, "DB Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"status":      "success",
		"new_status":  string(req.Status),
		"property_id": req.PropertyID,
	}

	if tx != nil {
		response["tx_hash"] = tx.Hash().Hex()
		response["message"] = "Transaction submitted. DB updated."
	}

	render.JSON(w, r, response)
}

// ---------------------------------------------------------
// HELPER FUNCTIONS
// ---------------------------------------------------------

type PropertyPayload struct {
	OwnerAddress string `json:"owner_address"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	Valuation    int64  `json:"valuation"`
	TokenSupply  int64  `json:"token_supply"`
}

func (handler *RequestHandler) parseCreatePropertyRequest(r *http.Request) (*PropertyPayload, []*multipart.FileHeader, error) {
	if err := r.ParseMultipartForm(20 << 20); err != nil {
		return nil, nil, fmt.Errorf("file too large or invalid format")
	}

	payloadData := r.FormValue("data")
	if payloadData == "" {
		return nil, nil, fmt.Errorf("missing 'data' field")
	}

	var payload PropertyPayload
	if err := json.Unmarshal([]byte(payloadData), &payload); err != nil {
		return nil, nil, fmt.Errorf("invalid JSON: %v", err)
	}

	if payload.OwnerAddress == "" || payload.Name == "" || payload.Valuation <= 0 {
		return nil, nil, fmt.Errorf("invalid payload fields")
	}

	files := r.MultipartForm.File["files"]
	if len(files) == 0 {
		return nil, nil, fmt.Errorf("at least one file is required")
	}

	return &payload, files, nil
}

func (handler *RequestHandler) processPropertyFiles(files []*multipart.FileHeader) ([]models.PropertyDocument, string, error) {
	var dbDocs []models.PropertyDocument
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

		dbDocs = append(dbDocs, models.PropertyDocument{
			ID:         uuid.New(),
			FileUrl:    fullUrl,
			FileHash:   hash, // Store IPFS hash
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

func inferDocType(filename string) string {
	lowerName := strings.ToLower(filename)

	switch {
	case strings.Contains(lowerName, "deed"):
		return "Deed"
	case strings.Contains(lowerName, "inspection"):
		return "Inspection Report"
	case strings.Contains(lowerName, "valuation"), strings.Contains(lowerName, "appraisal"):
		return "Valuation Report"
	case strings.Contains(lowerName, "image"), strings.Contains(lowerName, "img"):
		return "Photo"
	case strings.Contains(lowerName, "contract"), strings.Contains(lowerName, "agreement"):
		return "Legal Contract"
	default:
		return "General Document"
	}
}

func (handler *RequestHandler) createPropertyOnChain(p *PropertyPayload, dataHash string) (*blockchain.PropertyCreationResult, error) {
	if handler.chain == nil {
		return nil, fmt.Errorf("blockchain service not available")
	}

	return handler.chain.CreateProperty(
		p.OwnerAddress,
		p.Name,
		p.Symbol,
		dataHash,
		p.Valuation,
		p.TokenSupply,
	)
}
