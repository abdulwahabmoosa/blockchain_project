// api package - handles HTTP requests and responses
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

// Login - authenticate user and return JWT token
// POST /login - expects email and password in request body
func (handler *RequestHandler) Login(w http.ResponseWriter, r *http.Request) {
	log.Printf("Debug Login: Handler called - Method: %s, Path: %s", r.Method, r.URL.Path)
	
	// Ensure we always send a response, even on panic
	defer func() {
		if rec := recover(); rec != nil {
			log.Printf("Error Login: Panic recovered: %v", rec)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}()

	log.Printf("Debug Login: Starting request processing")

	// Safety check for validate
	if validate == nil {
		log.Printf("Error Login: CRITICAL - validate is nil!")
		http.Error(w, "Server configuration error", http.StatusInternalServerError)
		return
	}
	
	var creds auth.LoginCredentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		log.Printf("Error Login: Failed to decode request body: %v", err)
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	log.Printf("Debug Login: Decoded credentials - Email: %s", creds.Email)
	
	// Validate credentials - wrap in recover to catch any panics
	validationErr := func() (err error) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("Error Login: Validation panic: %v", rec)
				err = fmt.Errorf("validation panic: %v", rec)
			}
		}()
		return validate.Struct(creds)
	}()
	
	if validationErr != nil {
		log.Printf("Error Login: Validation failed: %v", validationErr)
		http.Error(w, "Validation Error: "+validationErr.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Debug Login: Validation passed")

	// Check if database handler is available
	if handler.db == nil {
		log.Printf("Error Login: CRITICAL - Database handler is nil!")
		http.Error(w, "Database connection not available", http.StatusInternalServerError)
		return
	}

	user, err := handler.db.GetUserByCredentials(creds)
	if err != nil {
		log.Printf("Error Login: Database error for email %s: %v", creds.Email, err)
		// Check if it's a "record not found" error (user doesn't exist) or password mismatch
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Error Login: User not found for email %s", creds.Email)
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		} else {
			// Other database errors (connection issues, etc.) or password mismatch
			log.Printf("Error Login: Database query failed or password mismatch: %v", err)
			// Password mismatch from bcrypt will also come here, treat as invalid credentials
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		}
		return
	}

	log.Printf("Info Login: User retrieved - ID: %s, Email: %s, Role: %s, ApprovalStatus: %s",
		user.ID.String(), user.Email, string(user.Role), string(user.ApprovalStatus))

	token, err := auth.GenerateToken(user.ID, user.Email)
	if err != nil {
		log.Printf("Error Login: Failed to generate token: %v", err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	// Create a safe user response (exclude password hashs)
	// Handle potential empty/null values safely
	approvalStatus := string(user.ApprovalStatus)
	if approvalStatus == "" {
		approvalStatus = "pending" // Default value
		log.Printf("Warning Login: ApprovalStatus was empty, defaulting to 'pending'")
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

	log.Printf("Info Login: Attempting to serialize response for user %s", user.Email)

	// Wrap render.JSON in recover to catch any serialization panics
	func() {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("Error Login: JSON serialization panic: %v", rec)
				// Try to send a basic error response if we haven't written headers yet
				if w.Header().Get("Content-Type") == "" {
					http.Error(w, fmt.Sprintf("Response serialization error: %v", rec), http.StatusInternalServerError)
				}
			}
		}()
		render.JSON(w, r, response)
	}()

	log.Printf("Success Login: Successfully authenticated user %s", user.Email)
}

// RegisterUser - create new user account
// POST /register - expects email, password, name, wallet_address
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
