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

func (db DBHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	var user models.User
	vars := mux.Vars(r)
	userId, idOk := vars["id"]

	if !idOk {
		http.Error(w, "User ID not provided", http.StatusBadRequest)
		return
	}

	result := db.DB.First(&user, "id = ?", userId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, fmt.Sprintf("User with ID: %s not found", userId), http.StatusNotFound)
		return
	}

	if result.Error != nil {
		log.Println("Error getting the user", result.Error)
	}

	json.NewEncoder(w).Encode(user)
}
