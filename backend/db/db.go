package db

import (
	"backend/auth"
	"backend/db/models"
	"context"
	"errors"
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
	            CREATE TYPE user_role AS ENUM ('admin', 'investor', 'owner');
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
	return err
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

func (db *Database) CreateProperty(prop models.Property) error {
	return gorm.G[models.Property](db.db).Create(db.ctx, &prop)
}

func (db *Database) CreateRevenueDistribution(rev models.RevenueDistribution) error {
	return gorm.G[models.RevenueDistribution](db.db).Create(db.ctx, &rev)
}

func (db *Database) CreateUser(details auth.RegisterUserPayload) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(details.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := models.User{
		Email:        details.Email,
		Name:         details.Name,
		PasswordHash: string(hash),
	}

	return gorm.G[models.User](db.db).Create(db.ctx, &newUser)
}

func (db *Database) GetUserById(id int) (models.User, error) {
	return gorm.G[models.User](db.db).Where("id = ?", id).First(db.ctx)
}

func (db *Database) UserExists(email string) (bool, error) {
	_, err := gorm.G[models.User](db.db).Where("email = ?", email).First(db.ctx)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (db *Database) GetPropertyByTokenAddress(addr string) (result models.Property, err error) {
	result, err = gorm.G[models.Property](db.db).Where("onchain_token_address = ?", addr).First(db.ctx)
	return
}

func (db *Database) GetUserByCredentials(creds auth.LoginCredentials) (user models.User, err error) {
	user, err = gorm.G[models.User](db.db).
		Where("email = ?", creds.Email).
		First(db.ctx)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(creds.Password),
	)
	if err != nil {
		return
	}

	return
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
