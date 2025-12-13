package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserRole represents the possible roles for a user.
type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleUser  UserRole = "user"
)

type User struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`                    // Maps to id (UUID/INT), using UUID for consistency
	WalletAddress string    `gorm:"type:varchar(100);unique;not null;index"` // MUST be unique, links to on-chain addresses
	Email         string    `gorm:"type:varchar(255)"`
	Name          string    `gorm:"type:varchar(255)"`
	PasswordHash  string    `gorm:"not null"`
	Role          UserRole  `gorm:"type:user_role;default:'user'"` // Maps to role ENUM
	IsApproved    bool      `gorm:"default:false"`                 // Synced from ApprovalService contract
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
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
	OnchainAssetAddress string         `gorm:"type:varchar(100);not null"`            // PropertyAsset (ERC721) contract address
	OnchainTokenAddress string         `gorm:"type:varchar(100);not null"`            // PropertyToken (ERC20) contract address
	OwnerWallet         string         `gorm:"type:varchar(100);not null;index"`      // FK relationship with User (Owner)
	MetadataHash        string         `gorm:"type:varchar(255);not null"`            // IPFS/document hash
	Valuation           float64        `gorm:"type:decimal"`                          // DECIMAL, using float64 or string for precision
	Status              PropertyStatus `gorm:"type:property_status;default:'Active'"` // Synced from blockchain
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

type Transaction struct {
	ID            uuid.UUID       `gorm:"type:uuid;primaryKey"`
	WalletAddress string          `gorm:"type:varchar(100);not null;index"` // Who performed the action
	Type          TransactionType `gorm:"type:transaction_type;not null"`
	TxHash        string          `gorm:"type:varchar(100)"` // Blockchain transaction hash
	Details       string          `gorm:"type:json"`         // JSON column for extra metadata
	CreatedAt     time.Time
}
