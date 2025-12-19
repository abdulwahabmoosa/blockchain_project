package models

import (
	"time"

	"github.com/google/uuid"
)

// UserRole represents the possible roles for a user.
type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleUser  UserRole = "user"
)

// ApprovalStatus represents the approval status for users and properties.
type ApprovalStatus string

const (
	ApprovalPending  ApprovalStatus = "pending"
	ApprovalApproved ApprovalStatus = "approved"
	ApprovalRejected ApprovalStatus = "rejected"
)

type User struct {
	ID             uuid.UUID      `json:"ID" gorm:"type:uuid;primaryKey"`
	WalletAddress  string         `json:"WalletAddress" gorm:"type:varchar(100);unique;not null;index"`
	Email          string         `json:"Email" gorm:"type:varchar(255)"`
	Name           string         `json:"Name" gorm:"type:varchar(255)"`
	PasswordHash   string         `json:"-" gorm:"not null"` // Exclude from JSON
	Role           UserRole       `json:"Role" gorm:"type:user_role;default:'user'"`
	ApprovalStatus ApprovalStatus `json:"approval_status" gorm:"type:approval_status;default:'pending'"`
	CreatedAt      time.Time      `json:"CreatedAt"`
	UpdatedAt      time.Time      `json:"UpdatedAt"`
}

// PropertyStatus represents the status ENUM for a property.
type PropertyStatus string

const (
	StatusActive   PropertyStatus = "Active"
	StatusPaused   PropertyStatus = "Paused"
	StatusDisputed PropertyStatus = "Disputed"
	StatusClosed   PropertyStatus = "Closed"
)

type Property struct {
	ID                  uuid.UUID      `gorm:"type:uuid;primaryKey"`                  // Maps to id (UUID/INT)
	Name                string         `gorm:"type:varchar(255)"`                      // Property name
	OnchainAssetAddress string         `gorm:"type:varchar(100);not null"`            // PropertyAsset (ERC721) contract address
	OnchainTokenAddress string         `gorm:"type:varchar(100);not null"`            // PropertyToken (ERC20) contract address
	OwnerWallet         string         `gorm:"type:varchar(100);not null;index"`      // FK relationship with User (Owner)
	MetadataHash        string         `gorm:"type:varchar(255);not null"`            // IPFS/document hash
	Valuation           float64        `gorm:"type:decimal"`                          // DECIMAL, using float64 or string for precision
	Status              PropertyStatus `gorm:"type:property_status;default:'Active'"` // Synced from blockchain
	TxHash              string         `gorm:"type:varchar(100)"`                     // Transaction hash of property creation
	CreatedAt           time.Time
	// Relationships
	Documents []PropertyDocument    `gorm:"foreignKey:PropertyID"`
	Revenues  []RevenueDistribution `gorm:"foreignKey:PropertyID"`
}

type PropertyDocument struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	PropertyID uuid.UUID `gorm:"type:uuid;not null;index"` // FK(properties.id)
	FileUrl    string    `gorm:"type:text;not null"`
	FileHash   string    `gorm:"type:varchar(255);not null"` // Hash must match on-chain hash
	Name       string    `gorm:"type:varchar(255);not null"` // File name
	Type       string    `gorm:"type:varchar(100);not null"` // Document type
	UploadedAt time.Time
	// Relationships
	Property Property `gorm:"foreignKey:PropertyID"`
}

type RevenueDistribution struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey"`
	PropertyID       uuid.UUID `gorm:"type:uuid;not null;index"`   // FK(properties.id)
	SnapshotID       int32     `gorm:"type:int;not null"`          // Matches on-chain snapshot ID
	StablecoinTxHash string    `gorm:"type:varchar(100);not null"` // Deposit transaction hash
	TotalAmount      string    `gorm:"type:decimal;not null"`      // Using string for high-precision DECIMAL
	CreatedAt        time.Time
	// Relationships
	Property Property       `gorm:"foreignKey:PropertyID"`
	Claims   []RevenueClaim `gorm:"foreignKey:RevenueDistributionID"`
}

type RevenueClaim struct {
	ID                    uuid.UUID `gorm:"type:uuid;primaryKey"`
	RevenueDistributionID uuid.UUID `gorm:"type:uuid;not null;index"`         // FK(revenue_distributions.id)
	WalletAddress         string    `gorm:"type:varchar(100);not null;index"` // Investor wallet address
	Amount                string    `gorm:"type:decimal;not null"`            // Using string for claimed revenue DECIMAL
	TxHash                string    `gorm:"type:varchar(100);not null"`       // Blockchain transaction hash
	ClaimedAt             time.Time
	// Relationships
	Distribution RevenueDistribution `gorm:"foreignKey:RevenueDistributionID"`
}

// TransactionType represents the actions recorded in the audit log.
type TransactionType string

const (
	TxTypeApproveUser    TransactionType = "approve_user"
	TxTypeCreateProperty TransactionType = "create_property"
	TxTypeDepositRevenue TransactionType = "deposit_revenue"
	TxTypeClaimRevenue   TransactionType = "claim_revenue"
	TxTypeTransferToken  TransactionType = "transfer_token"
)

// TransactionStatus represents the status of blockchain transactions.
type TransactionStatus string

const (
	TxStatusPending   TransactionStatus = "pending"
	TxStatusConfirmed TransactionStatus = "confirmed"
	TxStatusFailed    TransactionStatus = "failed"
)

type Transaction struct {
	ID            uuid.UUID         `gorm:"type:uuid;primaryKey"`
	WalletAddress string            `gorm:"type:varchar(100);not null;index"` // Who performed the action
	Type          TransactionType   `gorm:"type:transaction_type;not null"`
	Status        TransactionStatus `gorm:"type:transaction_status;default:'pending'"` // Transaction status
	TxHash        string            `gorm:"type:varchar(100)"`                         // Blockchain transaction hash
	Details       string            `gorm:"type:json"`                                 // JSON column for extra metadata
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
