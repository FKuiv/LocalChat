package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/FKuiv/LocalChat/pkg/models"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func (db dbHandler) SaveMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var message models.MessageRequest
	err := json.NewDecoder(r.Body).Decode(&message)

	if err != nil {
		log.Println("Error in /message POST", err)
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	messageId, messageIdErr := gonanoid.New()

	if messageIdErr != nil {
		log.Println("Error creating message ID", messageIdErr)
		http.Error(w, "Error creating ID for message", http.StatusInternalServerError)
		return
	}

	newMessage := &models.Message{ID: messageId, Content: message.Content, UserID: message.UserID, GroupID: message.GroupID}
	result := db.DB.Create(newMessage)

	if result.Error != nil {
		log.Println("error saving message")
		http.Error(w, "Error creating message", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(newMessage)
}
