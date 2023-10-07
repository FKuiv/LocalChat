package models

import "time"

type User struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"CreatedAt"`
}

type UserCreateReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
