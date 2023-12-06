package repos

import (
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

// Repositories contains all the repo structs
type Repositories struct {
	UserRepo    *UserRepo
	GroupRepo   *GroupRepo
	MessageRepo *MessageRepo
}

// InitRepositories should be called in main.go
func InitRepositories(db *gorm.DB, minio *minio.Client) *Repositories {
	userRepo := NewUserRepo(db, minio)
	groupRepo := NewGroupRepo(db, minio)
	messageRepo := NewMessageRepo(db, minio)
	return &Repositories{
		UserRepo:    userRepo,
		GroupRepo:   groupRepo,
		MessageRepo: messageRepo,
	}
}
