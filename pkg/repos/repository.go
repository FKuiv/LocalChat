package repos

import (
	"gorm.io/gorm"
)

// Repositories contains all the repo structs
type Repositories struct {
	UserRepo    *UserRepo
	GroupRepo   *GroupRepo
	MessageRepo *MessageRepo
}

// InitRepositories should be called in main.go
func InitRepositories(db *gorm.DB) *Repositories {
	userRepo := NewUserRepo(db)
	groupRepo := NewGroupRepo(db)
	messageRepo := NewMessageRepo(db)
	return &Repositories{
		UserRepo:    userRepo,
		GroupRepo:   groupRepo,
		MessageRepo: messageRepo,
	}
}
