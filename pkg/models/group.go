package models

import "time"

type Group struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Users     []*User   `gorm:"many2many:user_groups;" json:"users"`
	Messages  []Message `json:"messages"` // Every group can have a lot of messages
	Admins    []string  `gorm:"type:text" json:"admins"`
}

type GroupRequest struct {
	Name    string   `json:"name"`
	UserIDs []string `json:"user_ids"`
	Admins  []string `json:"admins"`
}
