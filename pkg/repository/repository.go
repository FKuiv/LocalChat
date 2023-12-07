package repository

import (
	"github.com/FKuiv/LocalChat/pkg/websocket"
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
func InitRepositories(db *gorm.DB, minio *minio.Client, hub *websocket.Hub) *Repositories {
	userRepo := NewUserRepo(db, minio, hub)
	groupRepo := NewGroupRepo(db, minio, hub)
	messageRepo := NewMessageRepo(db, minio, hub)
	return &Repositories{
		UserRepo:    userRepo,
		GroupRepo:   groupRepo,
		MessageRepo: messageRepo,
	}
}
