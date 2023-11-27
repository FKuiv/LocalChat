package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/FKuiv/LocalChat/pkg/models"
)

func (db DBHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	result := db.DB.Find(&users)

	if result.Error != nil {
		http.Error(w, "There was an error getting users", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}
