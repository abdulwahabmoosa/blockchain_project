package db

import (
	"backend/auth"
	"backend/db/models"
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db  *gorm.DB
	ctx context.Context
}

func (db *Database) createEnums() error {
	sql := `
        DO $$
	    BEGIN
	        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN
	            CREATE TYPE user_role AS ENUM ('admin', 'user');
			END IF;
	        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'property_status') THEN
	            CREATE TYPE property_status AS ENUM ('Active', 'Paused', 'Disputed', 'Closed');
            END IF;
            IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'approval_status') THEN
                CREATE TYPE approval_status AS ENUM ('pending', 'approved', 'rejected');
            END IF;
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transaction_status') THEN
                CREATE TYPE transaction_status AS ENUM ('pending', 'confirmed', 'failed');
            END IF;
            IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transaction_type') THEN
            	CREATE TYPE transaction_type AS ENUM ('approve_user', 'create_property', 'deposit_revenue', 'claim_revenue', 'transfer_token');
            END IF;
        END
        $$;
    `
	if err := db.db.Exec(sql).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) seedAdmin() error {
	// 1. Configuration (Env or Default)
	email := os.Getenv("ADMIN_EMAIL")
	if email == "" {
		email = "admin@admin.com"
	}
	password := os.Getenv("ADMIN_PASSWORD")
	if password == "" {
		password = "password"
	}
	wallet := os.Getenv("ADMIN_WALLET")
	if wallet == "" {
		wallet = "0x98E5a749E25C56e19C28008505DF75aFf4988049" // Admin wallet address
	}

	// 2. Generate Hash in Go
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 3. Construct Raw SQL Block
	// This uses the exact same 'DO $$' style as createEnums
	// We inject values using fmt.Sprintf for the seed data
	sql := fmt.Sprintf(`
		DO $$
		DECLARE
			target_email TEXT := '%s';
			target_wallet TEXT := '%s';
			target_hash TEXT := '%s';
			new_id uuid := '%s';
		BEGIN
			-- Check if user exists by email
			IF EXISTS (SELECT 1 FROM users WHERE email = target_email) THEN
				-- Update existing user to be admin
				UPDATE users
				SET role = 'admin', approval_status = 'approved', password_hash = target_hash
				WHERE email = target_email;
			ELSE
				-- Insert new admin user
				INSERT INTO users (id, email, name, wallet_address, password_hash, role, approval_status, created_at, updated_at)
				VALUES (new_id, target_email, 'System Admin', target_wallet, target_hash, 'admin', 'approved', NOW(), NOW());
			END IF;
		END
		$$;
	`, email, wallet, hash, uuid.New())

	// 4. Execute Raw SQL
	return db.db.Exec(sql).Error
}

func (db *Database) migrate() error {
	if err := db.createEnums(); err != nil {
		return fmt.Errorf("failed to create enums: %w", err)
	}

	log.Println("Info: Running database migrations...")
	err := db.db.AutoMigrate(
		&models.User{},
		&models.Property{},
		&models.PropertyDocument{},
		&models.RevenueDistribution{},
		&models.RevenueClaim{},
		&models.Transaction{},
		&models.PropertyUploadRequest{},
		&models.PropertyUploadRequestDocument{},
		&models.TokenPurchase{},
	)

	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	log.Println("Success: Database migrations completed successfully")
	return db.seedAdmin()
}

func NewDatabase() (db *Database, err error) {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return nil, errors.New("DB_URL environment variable is required")
	}

	gormDB, err := gorm.Open(
		postgres.Open(dbURL),
	)

	if err != nil {
		return
	}

	db = &Database{db: gormDB, ctx: context.Background()}
	if err = db.migrate(); err != nil {
		return
	}
	return
}

func (db *Database) CreateUser(details auth.RegisterUserPayload) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(details.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	newUser := models.User{
		ID:            uuid.New(),
		Email:         details.Email,
		Name:          details.Name,
		WalletAddress: details.WalletAddress,
		PasswordHash:  string(hash),
		Role:          models.RoleUser,
	}

	return gorm.G[models.User](db.db).Create(db.ctx, &newUser)
}

func (db *Database) GetUserById(id string) (models.User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return models.User{}, err
	}
	return gorm.G[models.User](db.db).Where("id = ?", uid).First(db.ctx)
}

func (db *Database) GetUserByWallet(wallet string) (models.User, error) {
	var user models.User
	result := db.db.WithContext(db.ctx).Where("wallet_address = ?", wallet).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (db *Database) GetAllUsers() (result []models.User, err error) {
	if db.db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	result, err = gorm.G[models.User](db.db).Find(db.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %w", err)
	}
	// Return empty slice instead of nil if no users found
	if result == nil {
		result = []models.User{}
	}
	return result, nil
}

func (db *Database) UpdateUserApproval(wallet string, status models.ApprovalStatus) error {
	var approved models.ApprovalStatus
	switch status {
	case models.ApprovalApproved:
		approved = models.ApprovalApproved
	case models.ApprovalRejected:
		approved = models.ApprovalRejected
	default:
		return errors.New("invalid approval status")
	}

	_, err := gorm.G[models.User](db.db).
		Where("wallet_address = ?", wallet).
		Update(db.ctx, "approval_status", approved)
	return err
}

// --- Property Methods ---

func (db *Database) CreateProperty(prop models.Property) error {
	return gorm.G[models.Property](db.db).Create(db.ctx, &prop)
}

func (db *Database) CreatePropertyDocument(doc models.PropertyDocument) error {
	return gorm.G[models.PropertyDocument](db.db).Create(db.ctx, &doc)
}

func (db *Database) GetAllProperties() (result []models.Property, err error) {
	result, err = gorm.G[models.Property](db.db).Where("status = ?", models.StatusActive).Find(db.ctx)
	return
}

func (db *Database) GetPropertyByID(id string) (result models.Property, err error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return
	}
	result, err = gorm.G[models.Property](db.db).Where("id = ?", uid).First(db.ctx)
	return
}

func (db *Database) GetPropertyByTokenAddress(addr string) (result models.Property, err error) {
	result, err = gorm.G[models.Property](db.db).Where("onchain_token_address = ?", addr).First(db.ctx)
	return
}

func (db *Database) GetPropertyByAssetAddress(addr string) (result models.Property, err error) {
	result, err = gorm.G[models.Property](db.db).Where("onchain_asset_address = ?", addr).First(db.ctx)
	return
}

func (db *Database) UpdatePropertyApproval(propertyID string, status models.ApprovalStatus) error {
	uid, err := uuid.Parse(propertyID)
	if err != nil {
		return err
	}

	// Map ApprovalStatus to PropertyStatus
	var propertyStatus models.PropertyStatus
	switch status {
	case models.ApprovalApproved:
		propertyStatus = models.StatusActive
	case models.ApprovalRejected:
		propertyStatus = models.StatusClosed
	default:
		return errors.New("invalid approval status")
	}

	_, err = gorm.G[models.Property](db.db).
		Where("id = ?", uid).
		Update(db.ctx, "status", propertyStatus)
	return err
}

// --- Revenue Methods ---

func (db *Database) CreateRevenueDistribution(rev models.RevenueDistribution) error {
	return gorm.G[models.RevenueDistribution](db.db).Create(db.ctx, &rev)
}

func (db *Database) CreateRevenueClaim(claim models.RevenueClaim) error {
	return gorm.G[models.RevenueClaim](db.db).Create(db.ctx, &claim)
}

func (db *Database) GetRevenueDistributionBySnapshot(propertyID uuid.UUID, snapshotID int32) (models.RevenueDistribution, error) {
	return gorm.G[models.RevenueDistribution](db.db).
		Where("property_id = ? AND snapshot_id = ?", propertyID, snapshotID).
		First(db.ctx)
}

// --- Transaction Methods ---

func (db *Database) CreateTransaction(tx models.Transaction) error {
	return gorm.G[models.Transaction](db.db).Create(db.ctx, &tx)
}

func (db *Database) UpdateTransactionStatus(txHash string, status models.TransactionStatus) error {
	_, err := gorm.G[models.Transaction](db.db).
		Where("tx_hash = ?", txHash).
		Update(db.ctx, "status", status)
	return err
}

func (db *Database) GetTransactionByHash(txHash string) (models.Transaction, error) {
	return gorm.G[models.Transaction](db.db).Where("tx_hash = ?", txHash).First(db.ctx)
}

// --- Auth Helpers ---

func (db *Database) UserExists(email string) (bool, error) {
	_, err := gorm.G[models.User](db.db).Where("email = ?", email).First(db.ctx)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func (db *Database) GetUserByCredentials(creds auth.LoginCredentials) (user models.User, err error) {
	if db.db == nil {
		return user, fmt.Errorf("database connection is nil")
	}
	
	user, err = gorm.G[models.User](db.db).Where("email = ?", creds.Email).First(db.ctx)
	if err != nil {
		// Return the error as-is so the caller can distinguish between "not found" and other errors
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(creds.Password))
	return user, err
}

func (db *Database) UpdateUserInfo(userID string, name string, email string) error {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return err
	}

	user := models.User{}
	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = email
	}

	_, err = gorm.G[models.User](db.db).Where("id = ?", uid).Updates(db.ctx, user)
	return err
}

func (db *Database) DeleteUser(userID string) error {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return err
	}

	_, err = gorm.G[models.User](db.db.Unscoped()).Where("id = ?", uid).Delete(db.ctx)
	return err
}

func (db *Database) UpdatePassword(userID string, newHash string) error {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return err
	}

	_, err = gorm.G[models.User](db.db).Where("id = ?", uid).Update(db.ctx, "password_hash", newHash)
	return err
}

// --- Property Upload Request Methods ---

func (db *Database) CreatePropertyUploadRequest(request models.PropertyUploadRequest) error {
	return gorm.G[models.PropertyUploadRequest](db.db).Create(db.ctx, &request)
}

func (db *Database) CreatePropertyUploadRequestDocument(doc models.PropertyUploadRequestDocument) error {
	return gorm.G[models.PropertyUploadRequestDocument](db.db).Create(db.ctx, &doc)
}

func (db *Database) GetPropertyUploadRequests() (result []models.PropertyUploadRequest, err error) {
	result, err = gorm.G[models.PropertyUploadRequest](db.db).Order("created_at DESC").Find(db.ctx)
	return
}

func (db *Database) GetPropertyUploadRequestsByUser(walletAddress string) (result []models.PropertyUploadRequest, err error) {
	result, err = gorm.G[models.PropertyUploadRequest](db.db).
		Where("wallet_address = ?", walletAddress).
		Order("created_at DESC").
		Find(db.ctx)
	return
}

func (db *Database) GetPropertyUploadRequestByID(id string) (result models.PropertyUploadRequest, err error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return
	}
	result, err = gorm.G[models.PropertyUploadRequest](db.db).Where("id = ?", uid).First(db.ctx)
	return
}

func (db *Database) UpdatePropertyUploadRequestStatus(id string, status models.ApprovalStatus, reason string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	updates := map[string]interface{}{
		"status": status,
	}
	if reason != "" {
		updates["rejection_reason"] = reason
	}

	result := db.db.WithContext(db.ctx).
		Model(&models.PropertyUploadRequest{}).
		Where("id = ?", uid).
		Updates(updates)
	return result.Error
}

// --- Token Purchase Methods ---

func (db *Database) CreateTokenPurchase(purchase models.TokenPurchase) error {
	return gorm.G[models.TokenPurchase](db.db).Create(db.ctx, &purchase)
}

func (db *Database) GetTokenPurchasesByProperty(propertyID string) (result []models.TokenPurchase, err error) {
	uid, err := uuid.Parse(propertyID)
	if err != nil {
		return nil, err
	}
	result, err = gorm.G[models.TokenPurchase](db.db).
		Where("property_id = ?", uid).
		Order("created_at DESC").
		Find(db.ctx)
	return
}

func (db *Database) GetTokenStatsByProperty(propertyID string) (totalSold float64, err error) {
	uid, err := uuid.Parse(propertyID)
	if err != nil {
		return 0, err
	}

	var result struct {
		TotalSold float64
	}

	// Use raw SQL to sum the amount column (stored as decimal/string)
	// PostgreSQL requires casting string to numeric for SUMM
	err = db.db.WithContext(db.ctx).
		Model(&models.TokenPurchase{}).
		Select("COALESCE(SUM(amount::numeric), 0) as total_sold").
		Where("property_id = ?", uid).
		Scan(&result).Error

	if err != nil {
		return 0, err
	}

	return result.TotalSold, nil
}

func (db *Database) GetTokenPurchasesByBuyer(buyerWallet string) (result []models.TokenPurchase, err error) {
	result, err = gorm.G[models.TokenPurchase](db.db).
		Where("buyer_wallet = ?", buyerWallet).
		Order("created_at DESC").
		Find(db.ctx)
	return
}

// GetPendingTokenPurchasesByProperty gets all pending token purchases (where TokenTxHash = "pending") for a property
func (db *Database) GetPendingTokenPurchasesByProperty(propertyID string) (result []models.TokenPurchase, err error) {
	uid, err := uuid.Parse(propertyID)
	if err != nil {
		return nil, err
	}
	result, err = gorm.G[models.TokenPurchase](db.db).
		Where("property_id = ? AND token_tx_hash = ?", uid, "pending").
		Order("created_at DESC").
		Find(db.ctx)
	return
}

// UpdateTokenPurchaseTxHash updates the token transfer transaction hash for a purchase
func (db *Database) UpdateTokenPurchaseTxHash(purchaseID string, tokenTxHash string) error {
	uid, err := uuid.Parse(purchaseID)
	if err != nil {
		return err
	}
	_, err = gorm.G[models.TokenPurchase](db.db).
		Where("id = ?", uid).
		Update(db.ctx, "token_tx_hash", tokenTxHash)
	return err
}
