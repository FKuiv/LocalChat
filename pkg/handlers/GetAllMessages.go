package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/FKuiv/LocalChat/pkg/models"
)

func (db DBHandler) GetAllMessages(w http.ResponseWriter, r *http.Request) {
	var messages []models.Message

	result := db.DB.Find(&messages)

	if result.Error != nil {
		log.Println("Error getting the messsages", result.Error)
	}

	json.NewEncoder(w).Encode(messages)
}
