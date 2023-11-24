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

func (db DBHandler) UpdateGroup(w http.ResponseWriter, r *http.Request) {
	var newGroupInfo models.GroupRequest
	err := json.NewDecoder(r.Body).Decode(&newGroupInfo)
	if utils.DecodingErr(err, "/group", w) {
		return
	}

	vars := mux.Vars(r)
	groupId, idOk := vars["id"]
	var currentGroup models.Group

	if utils.MuxVarsNotProvided(idOk, groupId, "Group ID", w) {
		return
	}

	result := db.DB.Where("id = ?", groupId).First(&currentGroup)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, fmt.Sprintf("Group with ID: %s not found", groupId), http.StatusNotFound)
		return
	}

	if result.Error != nil {
		log.Println("Error getting the group", result.Error)
		http.Error(w, "Error getting the group", http.StatusInternalServerError)
		return
	}

	if newGroupInfo.Name != "" {
		currentGroup.Name = newGroupInfo.Name
	} else {
		http.Error(w, "Group name cannot be empty", http.StatusBadRequest)
		return
	}

	if len(newGroupInfo.UserIDs) != 0 {
		var users []*models.User

		for _, userId := range newGroupInfo.UserIDs {
			var user *models.User
			result := db.DB.First(&user, "id = ?", userId)

			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				log.Println("Error finding user", result.Error)
				http.Error(w, "Error finding user", http.StatusInternalServerError)
			} else {
				users = append(users, user)
			}

		}

		currentGroup.Users = users
	} else {
		http.Error(w, "A group cannot have 0 users", http.StatusBadRequest)
		return
	}

	if len(newGroupInfo.Admins) != 0 {
		currentGroup.Admins = newGroupInfo.Admins
	} else {
		http.Error(w, "Group cannot have 0 admins", http.StatusBadRequest)
		return
	}

	db.DB.Save(&currentGroup)
	json.NewEncoder(w).Encode(currentGroup)
}
