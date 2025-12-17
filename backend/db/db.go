package db

import (
	"backend/auth"
	"backend/db/models"
	"context"
	"errors"
	"fmt"
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
				SET role = 'admin', is_approved = true, password_hash = target_hash
				WHERE email = target_email;
			ELSE
				-- Insert new admin user
				INSERT INTO users (id, email, name, wallet_address, password_hash, role, is_approved, created_at, updated_at)
				VALUES (new_id, target_email, 'System Admin', target_wallet, target_hash, 'admin', true, NOW(), NOW());
			END IF;
		END
		$$;
	`, email, wallet, hash, uuid.New())

	// 4. Execute Raw SQL
	return db.db.Exec(sql).Error
}

func (db *Database) migrate() error {
	if err := db.createEnums(); err != nil {
		return err
	}

	err := db.db.AutoMigrate(
		&models.User{},
		&models.Property{},
		&models.RevenueDistribution{},
		&models.RevenueClaim{},
		&models.Transaction{},
	)

	if err != nil {
		return err
	}

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
		return err
	}

	newUser := models.User{
		ID:            uuid.New(),
		Email:         details.Email,
		Name:          details.Name,
		WalletAddress: details.WalletAddress,
		PasswordHash:  string(hash),
		Role:          models.RoleUser,
		IsApproved:    false, // Default to false until on-chain approval
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
	return gorm.G[models.User](db.db).Where("wallet_address = ?", wallet).First(db.ctx)
}

func (db *Database) GetAllUsers() (result []models.User, err error) {
	result, err = gorm.G[models.User](db.db).Find(db.ctx)
	return
}

func (db *Database) UpdateUserApproval(wallet string, approved bool) error {
	_, err := gorm.G[models.User](db.db).
		Where("wallet_address = ?", wallet).
		Update(db.ctx, "is_approved", approved)
	return err
}

// --- Property Methods ---

func (db *Database) CreateProperty(prop models.Property) error {
	return gorm.G[models.Property](db.db).Create(db.ctx, &prop)
}

func (db *Database) GetAllProperties() (result []models.Property, err error) {
	result, err = gorm.G[models.Property](db.db).Find(db.ctx)
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
	user, err = gorm.G[models.User](db.db).Where("email = ?", creds.Email).First(db.ctx)
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(creds.Password))
	return
}
