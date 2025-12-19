package api

import (
	"backend/auth"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/render"
	"gorm.io/gorm"
)

func (handler *RequestHandler) Login(w http.ResponseWriter, r *http.Request) {
	log.Printf("🔵 Login: Handler called - Method: %s, Path: %s", r.Method, r.URL.Path)
	
	// Ensure we always send a response, even on panic
	defer func() {
		if rec := recover(); rec != nil {
			log.Printf("❌ Login: Panic recovered: %v", rec)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}()

	log.Printf("🔵 Login: Starting request processing")
	
	// Safety check for validate
	if validate == nil {
		log.Printf("❌ Login: CRITICAL - validate is nil!")
		http.Error(w, "Server configuration error", http.StatusInternalServerError)
		return
	}
	
	var creds auth.LoginCredentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		log.Printf("❌ Login: Failed to decode request body: %v", err)
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	log.Printf("🔵 Login: Decoded credentials - Email: %s", creds.Email)
	
	// Validate credentials - wrap in recover to catch any panics
	validationErr := func() (err error) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("❌ Login: Validation panic: %v", rec)
				err = fmt.Errorf("validation panic: %v", rec)
			}
		}()
		return validate.Struct(creds)
	}()
	
	if validationErr != nil {
		log.Printf("❌ Login: Validation failed: %v", validationErr)
		http.Error(w, "Validation Error: "+validationErr.Error(), http.StatusBadRequest)
		return
	}
	
	log.Printf("🔵 Login: Validation passed")

	// Check if database handler is available
	if handler.db == nil {
		log.Printf("❌ Login: CRITICAL - Database handler is nil!")
		http.Error(w, "Database connection not available", http.StatusInternalServerError)
		return
	}

	user, err := handler.db.GetUserByCredentials(creds)
	if err != nil {
		log.Printf("❌ Login: Database error for email %s: %v", creds.Email, err)
		// Check if it's a "record not found" error (user doesn't exist) or password mismatch
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("❌ Login: User not found for email %s", creds.Email)
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		} else {
			// Other database errors (connection issues, etc.) or password mismatch
			log.Printf("❌ Login: Database query failed or password mismatch: %v", err)
			// Password mismatch from bcrypt will also come here, treat as invalid credentials
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		}
		return
	}

	log.Printf("🔍 Login: User retrieved - ID: %s, Email: %s, Role: %s, ApprovalStatus: %s", 
		user.ID.String(), user.Email, string(user.Role), string(user.ApprovalStatus))

	token, err := auth.GenerateToken(user.ID, user.Email)
	if err != nil {
		log.Printf("❌ Login: Failed to generate token: %v", err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	// Create a safe user response (exclude password hash)
	// Handle potential empty/null values safely
	approvalStatus := string(user.ApprovalStatus)
	if approvalStatus == "" {
		approvalStatus = "pending" // Default value
		log.Printf("⚠️ Login: ApprovalStatus was empty, defaulting to 'pending'")
	}

	userResponse := map[string]any{
		"ID":             user.ID.String(),
		"Email":           user.Email,
		"Name":            user.Name,
		"WalletAddress":  user.WalletAddress,
		"Role":            string(user.Role),
		"ApprovalStatus":  approvalStatus,
		"CreatedAt":       user.CreatedAt,
		"UpdatedAt":       user.UpdatedAt,
	}

	response := map[string]any{
		"token": token,
		"user":  userResponse,
	}

	log.Printf("🔍 Login: Attempting to serialize response for user %s", user.Email)
	
	// Wrap render.JSON in recover to catch any serialization panics
	func() {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("❌ Login: JSON serialization panic: %v", rec)
				// Try to send a basic error response if we haven't written headers yet
				if w.Header().Get("Content-Type") == "" {
					http.Error(w, fmt.Sprintf("Response serialization error: %v", rec), http.StatusInternalServerError)
				}
			}
		}()
		render.JSON(w, r, response)
	}()
	
	log.Printf("✅ Login: Successfully authenticated user %s", user.Email)
}

func (handler *RequestHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var userDetails auth.RegisterUserPayload
	if err := json.NewDecoder(r.Body).Decode(&userDetails); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

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
