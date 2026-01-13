package auth

type RegisterUserPayload struct {
	// eth_addr checks if it starts with 0x and is 42 chars long (hex)
	WalletAddress string `json:"wallet_address" binding:"required" validate:"required,eth_addr"`
	Email         string `json:"email" binding:"required" validate:"required,email"`
	Password      string `json:"password" binding:"required" validate:"required,min=8"`
	Name          string `json:"name" validate:"required,min=2"`
}
