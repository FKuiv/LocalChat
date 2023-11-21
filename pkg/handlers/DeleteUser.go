package handlers

import (
	"fmt"
	"net/http"

	"github.com/FKuiv/LocalChat/pkg/models"
	"github.com/FKuiv/LocalChat/pkg/utils"
)

func (db DBHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("UserId")

	var user models.User
	result := db.DB.Where("id = ?", userId).First(&user)

	if utils.ItemNotFound(result.Error, "User", w) {
		return
	}

	if err := db.DB.Delete(&user).Error; err != nil {
		fmt.Println("Error deleting user:", err)
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	// Find and delete empty groups
	var emptyGroups []models.Group
	db.DB.Joins("LEFT JOIN user_groups ON groups.id = user_groups.group_id").
		Group("groups.id").
		Having("COUNT(user_groups.group_id) = 0").
		Find(&emptyGroups)

	for _, group := range emptyGroups {
		if err := db.DB.Unscoped().Model(&group).Association("Messages").Unscoped().Clear(); err != nil {
			fmt.Println("Error deleting all the messages in a group", err)
			http.Error(w, "Failed to delete all group messages", http.StatusInternalServerError)
			return
		}
		if err := db.DB.Delete(&group).Error; err != nil {
			fmt.Println("Error deleting group:", err)
			http.Error(w, "Failed to delete group", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}
