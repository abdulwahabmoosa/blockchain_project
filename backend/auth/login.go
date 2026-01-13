package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtKey = []byte("7umbl2i2Zy6AR6HjVaE2jN2RBmtzAZXrY4CgHmIFls4=")

type LoginCredentials struct {
	Email    string `json:"email" binding:"required" validate:"required,email"`
	Password string `json:"password" binding:"required" validate:"required"`
}

type Claims struct {
	UserID    uuid.UUID `json:"userId"`
	UserEmail string    `json:"userEmail"`
	jwt.RegisteredClaims
}

type TokenKeyType int
type ClaimsKeyType int

var TokenKey TokenKeyType
var ClaimsKey ClaimsKeyType

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
