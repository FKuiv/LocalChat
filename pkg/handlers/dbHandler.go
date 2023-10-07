package handlers

import "gorm.io/gorm"

type dbHandler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) dbHandler {
	return dbHandler{db}
}
