package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/FKuiv/LocalChat/pkg/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func (db DBHandler) UpdateMessage(w http.ResponseWriter, r *http.Request) {
	var newMessageInfo models.UpdateMessage
	err := json.NewDecoder(r.Body).Decode(&newMessageInfo)

	if err != nil {
		log.Println("Error in /message PATCH", err)
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	messageId, idOk := vars["id"]
	var currentMessage models.Message

	if !idOk {
		http.Error(w, "Message ID not provided", http.StatusBadRequest)
		return
	}

	result := db.DB.First(&currentMessage, "id = ?", messageId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, fmt.Sprintf("Message with ID: %s not found", messageId), http.StatusNotFound)
		return
	}

	if result.Error != nil {
		log.Println("Error getting the message", result.Error)
	}

	if newMessageInfo.Content != "" {
		currentMessage.Content = newMessageInfo.Content
	} else {
		http.Error(w, "Message cannot have empty content. If you want to delete it, then use delete.", http.StatusBadRequest)
	}

	db.DB.Save(&currentMessage)
	json.NewEncoder(w).Encode(currentMessage)
}
