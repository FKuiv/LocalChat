package models

import (
	"time"
)

type User struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Messages  []Message `json:"messages"`
	Groups    []*Group  `gorm:"many2many:user_groups;" json:"groups"`
	Session   Session   `json:"session"`
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
