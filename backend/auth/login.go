// auth package - handles user authentication and JWT tokens
package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// jwtKey - secret key for signing JWT tokens
var jwtKey = []byte("7umbl2i2Zy6AR6HjVaE2jN2RBmtzAZXrY4CgHmIFls4=")

// LoginCredentials - struct for login request data
type LoginCredentials struct {
	Email    string `json:"email" binding:"required" validate:"required,email"`
	Password string `json:"password" binding:"required" validate:"required"`
}

// Claims - JWT token payload structure
type Claims struct {
	UserID    uuid.UUID `json:"userId"`
	UserEmail string    `json:"userEmail"`
	jwt.RegisteredClaims
}

// context key types for request context
type TokenKeyType int
type ClaimsKeyType int

var TokenKey TokenKeyType
var ClaimsKey ClaimsKeyType

// GenerateToken - create JWT token for authenticated user
func GenerateToken(id uuid.UUID, email string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		UserID:    id,
		UserEmail: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "access",
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
