package httpserver

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/FKuiv/LocalChat/pkg/models"
	"github.com/FKuiv/LocalChat/pkg/utils"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func (db dbHandler) registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userInfo models.UserCreateReq
	err := json.NewDecoder(r.Body).Decode(&userInfo)

	if err != nil {
		log.Println("Error in /register", err)
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	log.Println("POST request to /register, data:", userInfo)

	passwordHash, err := utils.HashPassword(userInfo.Password)

	if err != nil {
		log.Println("Error hashing the password", err)
		http.Error(w, "Problem hashing the password", http.StatusBadGateway)
	}

	userId, userIdErr := gonanoid.New()

	if userIdErr != nil {
		log.Println("error in creating ID", userIdErr)
	}

	newUser := models.User{ID: userId, Username: userInfo.Username, Password: passwordHash, CreatedAt: time.Now()}
	result := db.DB.Create(newUser)
	if result.Error != nil {
		log.Println("error in creating user", result.Error)
	}

	json.NewEncoder(w).Encode(newUser.ID)
}
