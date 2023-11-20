package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/FKuiv/LocalChat/pkg/models"
)

func (db DBHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	result := db.DB.Find(&users)

	if result.Error != nil {
		log.Println("Error getting the users", result.Error)
	}

	json.NewEncoder(w).Encode(users)
}
