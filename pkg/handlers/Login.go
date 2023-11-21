package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/FKuiv/LocalChat/pkg/models"
	"github.com/FKuiv/LocalChat/pkg/utils"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func (db DBHandler) Login(w http.ResponseWriter, r *http.Request) {
	var userInfo models.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userInfo)
	if utils.DecodingErr(err, "/login", w) {
		return
	}

	var currentUser models.User
	db.DB.Find(&currentUser, "username = ?", userInfo.Username)

	if !utils.CheckPasswordHash(userInfo.Password, currentUser.Password) {
		http.Error(w, "Wrong password", http.StatusBadRequest)
		return
	}

	sessionId, idErr := gonanoid.New()
	if utils.IDCreationErr(idErr, w) {
		return
	}
	newSession := &models.Session{ID: sessionId, UserID: currentUser.ID, ExpiresAt: time.Now().AddDate(0, 0, 7)}
	result := db.DB.Create(newSession)
	if utils.CreationErr(result.Error, w) {
		return
	}

	json.NewEncoder(w).Encode(newSession)
}
