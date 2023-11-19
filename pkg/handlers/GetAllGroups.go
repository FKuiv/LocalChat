package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/FKuiv/LocalChat/pkg/models"
)

func (db DBHandler) GetAllGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var groups []models.Group

	// A way to also return users
	// result := db.DB.Model(&models.Group{}).Preload("Messages").Find(&groups)

	result := db.DB.Find(&groups)

	if result.Error != nil {
		log.Println("Error getting groups", result.Error)
		http.Error(w, "Error getting the groups", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(groups)
}
