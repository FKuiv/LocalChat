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

func (db DBHandler) GetGroupById(w http.ResponseWriter, r *http.Request) {
	var group models.Group
	vars := mux.Vars(r)
	groupId, idOk := vars["id"]

	if utils.MuxVarsNotProvided(idOk, groupId, "Group ID", w) {
		return
	}

	result := db.DB.First(&group, "id = ?", groupId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, fmt.Sprintf("Group with ID: %s not found", groupId), http.StatusNotFound)
		return
	}

	if result.Error != nil {
		log.Println("Error getting the group", result.Error)
	}

	json.NewEncoder(w).Encode(group)
}
