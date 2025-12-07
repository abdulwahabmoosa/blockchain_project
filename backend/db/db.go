package db

import (
	"backend/auth"
	"backend/db/models"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db  *gorm.DB
	ctx context.Context
}

func NewDatabase() (db *Database, err error) {
	gormDB, err := gorm.Open(
		postgres.Open("postgres://blockchain-db:password@localhost:5432/db?sslmode=disable"),
	)

	if err != nil {
		return
	}

	db = &Database{db: gormDB, ctx: context.Background()}
	return
}

func (db *Database) CreateUser(details auth.UserDetails) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(details.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := models.User{
		Email:        details.Email,
		FirstName:    details.FirstName,
		LastName:     details.LastName,
		Phone:        details.Phone,
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
