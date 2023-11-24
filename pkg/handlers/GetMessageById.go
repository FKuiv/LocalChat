package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/FKuiv/LocalChat/pkg/models"
	"github.com/FKuiv/LocalChat/pkg/utils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func (db DBHandler) GetMessageById(w http.ResponseWriter, r *http.Request) {
	var message models.Message
	vars := mux.Vars(r)
	messageId, idOk := vars["id"]

	if utils.MuxVarsNotProvided(idOk, messageId, "Message ID", w) {
		return
	}

	result := db.DB.First(&message, "id = ?", messageId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, fmt.Sprintf("Message with ID: %s not found", messageId), http.StatusNotFound)
		return
	}

	if result.Error != nil {
		log.Println("Error getting the message", result.Error)
	}

	json.NewEncoder(w).Encode(message)
}
