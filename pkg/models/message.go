package models

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	UserID    string         `json:"user_id"`  // aka the Author of the message
	GroupID   string         `json:"group_id"` // to filter the message into the right chat
	Content   string         `json:"content"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UpdateMessage struct {
	Content string `json:"content"`
}
