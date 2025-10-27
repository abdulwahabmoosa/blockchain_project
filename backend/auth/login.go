package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v5"
)

// NOTE: generated using: openssl rand -base64 32
var jwtKey = []byte("7umbl2i2Zy6AR6HjVaE2jN2RBmtzAZXrY4CgHmIFls4=")

type LoginCredentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Claims struct {
	UserID int `json:"userId"`
	jwt.RegisteredClaims
}

type TokenKey int

var tokenKey TokenKey

func generateToken(creds LoginCredentials) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)

	claims := &Claims{
		UserID: 0,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds LoginCredentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	token, err := generateToken(creds)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	ctx := context.WithValue(r.Context(), tokenKey, token)
	token, ok := ctx.Value(tokenKey).(string)
	if !ok {
		status := http.StatusUnprocessableEntity
		http.Error(w, http.StatusText(status), status)
		return
	}
	render.JSON(w, r, map[string]string{"token": token})
}
