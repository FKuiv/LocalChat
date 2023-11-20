package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/FKuiv/LocalChat/pkg/models"
	"github.com/FKuiv/LocalChat/pkg/utils"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func (db DBHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userInfo models.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userInfo)

	if err != nil {
		log.Println("Error in /register", err)
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	passwordHash, err := utils.HashPassword(userInfo.Password)

	if err != nil {
		log.Println("Error hashing the password", err)
		http.Error(w, "Problem hashing the password", http.StatusInternalServerError)
	}

	userId, userIdErr := gonanoid.New()

	utils.IDErr(userIdErr, w)

	newUser := &models.User{ID: userId, Username: userInfo.Username, Password: passwordHash}
	result := db.DB.Create(newUser)
	utils.CreationErr(result.Error, w)
	// It is a hacky solution but GORM doesn't have an error type to check the unique key constraint so I am checking the substring in the error
	if result.Error != nil && strings.Contains(result.Error.Error(), "(SQLSTATE 23505)") {
		http.Error(w, fmt.Sprintf("Username %s is already taken", newUser.Username), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(newUser)
}
