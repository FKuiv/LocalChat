package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/FKuiv/LocalChat/pkg/models"
	"github.com/FKuiv/LocalChat/pkg/utils"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

func (db DBHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	var group models.GroupRequest
	err := json.NewDecoder(r.Body).Decode(&group)
	if utils.DecodingErr(err, "/group", w) {
		return
	}

	if len(group.Admins) == 0 || len(group.UserIDs) == 0 {
		http.Error(w, "There needs to be at least 1 admin and user in group", http.StatusBadRequest)
		return
	}

	groupId, groupIdErr := gonanoid.New()

	if utils.IDCreationErr(groupIdErr, w) {
		return
	}

	var users []*models.User

	for _, userId := range group.UserIDs {
		var user *models.User
		result := db.DB.First(&user, "id = ?", userId)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println("Error finding user", result.Error)
		} else {
			users = append(users, user)
		}

	}

	newGroup := &models.Group{ID: groupId, Name: group.Name, Users: users, Admins: group.Admins, IsDm: group.IsDm}
	result := db.DB.Create(newGroup)

	if utils.CreationErr(result.Error, w) {
		return
	}

	json.NewEncoder(w).Encode(newGroup)

}
