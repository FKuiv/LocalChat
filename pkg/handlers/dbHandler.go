package handlers

import "gorm.io/gorm"

type DBHandler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) DBHandler {
	return DBHandler{db}
}
