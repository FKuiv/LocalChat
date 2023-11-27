package repository

import (
	"gorm.io/gorm"
)

// Repositories contains all the repo structs
type Repositories struct {
	UserRepo *UserRepo
}

// InitRepositories should be called in main.go
func InitRepositories(db *gorm.DB) *Repositories {
	userRepo := NewUserRepo(db)
	return &Repositories{UserRepo: userRepo}
}
