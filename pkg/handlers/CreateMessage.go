package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/FKuiv/LocalChat/pkg/models"
	"github.com/FKuiv/LocalChat/pkg/utils"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func (db DBHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var message models.MessageRequest
	err := json.NewDecoder(r.Body).Decode(&message)
	if utils.DecodingErr(err, "/message", w) {
		return
	}

	userId := r.Header.Get("UserId")
	messageId, messageIdErr := gonanoid.New()

	if messageIdErr != nil {
		log.Println("Error creating message ID", messageIdErr)
		http.Error(w, "Error creating ID for message", http.StatusInternalServerError)
		return
	}

	newMessage := &models.Message{ID: messageId, Content: message.Content, UserID: userId, GroupID: message.GroupID}
	result := db.DB.Create(newMessage)

	if result.Error != nil {
		log.Println("error saving message")
		http.Error(w, "Error creating message", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(newMessage)
}
