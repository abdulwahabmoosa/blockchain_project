package auth

type RegisterUserPayload struct {
	WalletAddress string `json:"wallet_address" binding:"required"`
	Email         string `json:"email" binding:"required,email"`
	Password      string `json:"password" binding:"required,min=8"` // Plaintext password
	Name          string `json:"name"`
}
