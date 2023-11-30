package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/FKuiv/LocalChat/pkg/models"
)

func (db DBHandler) GetAllGroups(w http.ResponseWriter, r *http.Request) {
	var groups []models.Group

	result := db.DB.Find(&groups)

	if result.Error != nil {
		log.Println("Error getting groups", result.Error)
		http.Error(w, "Error getting the groups", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(groups)
}
