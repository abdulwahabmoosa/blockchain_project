package api

import (
	"backend/auth"
	"backend/db"
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type RequestHandler struct {
	db *db.Database
}

func NewRequestHandler(db *db.Database) *RequestHandler {
	return &RequestHandler{db}
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
