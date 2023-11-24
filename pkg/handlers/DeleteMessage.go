package handlers

import (
	"fmt"
	"net/http"

	"github.com/FKuiv/LocalChat/pkg/models"
	"github.com/FKuiv/LocalChat/pkg/utils"
	"github.com/gorilla/mux"
)

func (db DBHandler) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	messageId, idOk := mux.Vars(r)["id"]
	if utils.MuxVarsNotProvided(idOk, messageId, "Message ID", w) {
		return
	}

	userId := r.Header.Get("UserId")

	var message models.Message
	result := db.DB.Where("id = ?", messageId).First(&message)

	if utils.ItemNotFound(result.Error, "Message", w) {
		return
	}

	if message.UserID != userId && message.UserID != "" {
		http.Error(w, "User does not own this message, therefore cannot delete it", http.StatusForbidden)
		return
	}

	// Attempt to delete the message
	if err := db.DB.Delete(&message).Error; err != nil {
		fmt.Println("Error deleting message:", err)
		http.Error(w, "Failed to delete message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message deleted successfully"))
}
