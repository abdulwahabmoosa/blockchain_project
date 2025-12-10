package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `gorm:"not null;unique"`
	FirstName    string `gorm:"not null"`
	LastName     string `gorm:"not null"`
	Phone        string `gorm:"not null"`
	PasswordHash string `gorm:"not null"`
}
