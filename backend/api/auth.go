package api

import (
	"backend/auth"
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
)

func (handler *RequestHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds auth.LoginCredentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(creds); err != nil {
		http.Error(w, "Validation Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := handler.db.GetUserByCredentials(creds)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateToken(user.ID, user.Email)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

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
