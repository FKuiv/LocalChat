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

func (db DBHandler) UpdateGroup(w http.ResponseWriter, r *http.Request) {
	var newGroupInfo models.GroupRequest
	err := json.NewDecoder(r.Body).Decode(&newGroupInfo)

	if err != nil {
		log.Println("Error in /group PATCH", err)
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	groupId, idOk := vars["id"]
	var currentGroup models.Group

	if !idOk {
		http.Error(w, "Group ID not provided", http.StatusBadRequest)
		return
	}

	result := db.DB.First(&currentGroup, "id = ?", groupId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, fmt.Sprintf("Group with ID: %s not found", groupId), http.StatusNotFound)
		return
	}

	if result.Error != nil {
		log.Println("Error getting the group", result.Error)
	}

	if newGroupInfo.Name != "" {
		currentGroup.Name = newGroupInfo.Name
	} else {
		http.Error(w, "Group name cannot be empty", http.StatusBadRequest)
	}

	if len(newGroupInfo.UserIDs) != 0 {
		var users []*models.User

		for _, userId := range newGroupInfo.UserIDs {
			var user *models.User
			result := db.DB.First(&user, "id = ?", userId)

			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				log.Println("Error finding user", result.Error)
			} else {
				users = append(users, user)
			}

		}

		currentGroup.Users = users
	} else {
		http.Error(w, "A group cannot have 0 users", http.StatusBadRequest)
	}

	db.DB.Save(&currentGroup)
	json.NewEncoder(w).Encode(currentGroup)
}
