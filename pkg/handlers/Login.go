package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/FKuiv/LocalChat/pkg/models"
	"github.com/FKuiv/LocalChat/pkg/utils"
)

func (db dbHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userInfo models.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userInfo)

	if err != nil {
		log.Println("Error in /login", err)
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	if err != nil {
		log.Println("Error hashing the password", err)
		http.Error(w, "Problem hashing the password", http.StatusInternalServerError)
	}
	var currentUser models.User
	db.DB.Find(&currentUser, "username = ?", userInfo.Username)

	if !utils.CheckPasswordHash(userInfo.Password, currentUser.Password) {
		http.Error(w, "Wrong password", http.StatusBadRequest)
		return
	}

	// Make DB session
}
