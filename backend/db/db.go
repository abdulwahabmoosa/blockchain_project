package db

import (
	"backend/db/models"
	"context"

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

func (db *Database) GetUserById(id uint) (models.User, error) {
	return gorm.G[models.User](db.db).Where("id = ?", id).First(db.ctx)
}
