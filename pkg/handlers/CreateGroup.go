package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/FKuiv/LocalChat/pkg/models"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

func (db DBHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var group models.GroupRequest
	err := json.NewDecoder(r.Body).Decode(&group)

	if err != nil {
		log.Println("Error in /group POST", err)
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	groupId, groupIdErr := gonanoid.New()

	if groupIdErr != nil {
		log.Println("Error creating group ID", groupIdErr)
		http.Error(w, "Error creating group ID", http.StatusInternalServerError)
		return
	}

	var users []*models.User

	for i, userId := range group.UserIDs {
		log.Println("index", i)
		var user *models.User
		result := db.DB.First(&user, "id = ?", userId)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println("Error finding user", result.Error)
		} else {
			users = append(users, user)
		}

	}

	log.Println("Users array", users)

	newGroup := &models.Group{ID: groupId, Name: group.Name, Users: users}
	result := db.DB.Create(newGroup)

	if result.Error != nil {
		log.Println("Error creating the group", result.Error)
		http.Error(w, "Error creating the group", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newGroup)

}
