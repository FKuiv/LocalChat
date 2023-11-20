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

func (db DBHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var newUserInfo models.UserRequest
	err := json.NewDecoder(r.Body).Decode(&newUserInfo)

	if err != nil {
		log.Println("Error in /user PATCH", err)
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	userId, idOk := vars["id"]
	var currentUser models.User

	if !idOk {
		http.Error(w, "User ID not provided", http.StatusBadRequest)
		return
	}

	result := db.DB.First(&currentUser, "id = ?", userId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, fmt.Sprintf("User with ID: %s not found", userId), http.StatusNotFound)
		return
	}

	if result.Error != nil {
		log.Println("Error getting the user", result.Error)
	}

	if newUserInfo.Username != "" {
		usernameCheck := db.DB.Where("name = ?", newUserInfo.Username).First(&currentUser)

		if usernameCheck.RowsAffected == 1 {
			http.Error(w, "Username already exists", http.StatusBadRequest)
			return
		}

		currentUser.Username = newUserInfo.Username
	}

	if newUserInfo.Password != "" {
		passwordHash, err := utils.HashPassword(newUserInfo.Password)

		if err != nil {
			log.Println("Error hashing the password", err)
			http.Error(w, "Problem hashing the password", http.StatusBadGateway)
		}

		currentUser.Password = passwordHash
	}

	db.DB.Save(&currentUser)
	json.NewEncoder(w).Encode(currentUser)
}
