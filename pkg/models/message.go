package models

import "time"

type Message struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	UserID    string    `json:"user_id"` // aka the Author of the message
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateMessage struct {
	UserID  string `json:"user_id"`
	Content string `json:"content"`
}
