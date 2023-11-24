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
	result := db.DB.Preload("Groups").Preload("Session").Where("id = ?", userId).First(&user)

	if utils.ItemNotFound(result.Error, "User", w) {
		return
	}

	// Delete all connections to messages
	if err := db.DB.Model(&user).Association("Messages").Clear(); err != nil {
		fmt.Println("Error deleting assoctions of user messages", err)
		http.Error(w, "Failed to delete assoctions of messages", http.StatusInternalServerError)
		return
	}

	for _, group := range user.Groups {
		var newAdmins []string
		for _, adminId := range group.Admins {
			if adminId != userId {
				newAdmins = append(newAdmins, adminId)
			}
		}

		if len(group.Admins) == 1 && group.Admins[0] == userId {
			// Permanent delete
			if err := db.DB.Unscoped().Model(&group).Association("Messages").Unscoped().Clear(); err != nil {
				fmt.Println("Error deleting all the messages in a group", err)
				http.Error(w, "Failed to delete all group messages", http.StatusInternalServerError)
				return
			}

			if err := db.DB.Model(&group).Association("Users").Clear(); err != nil {
				fmt.Println("Error removing associations on group where user is only admine", err)
				http.Error(w, "Failed to delete group", http.StatusInternalServerError)
				return
			}

			if err := db.DB.Unscoped().Delete(&group).Error; err != nil {
				fmt.Println("Error deleting group where user is only admine", err)
				http.Error(w, "Failed to delete group", http.StatusInternalServerError)
				return
			}

		} else {
			if err := db.DB.Model(&user).Association("Groups").Delete(&group); err != nil {
				fmt.Println("Error removing association from group", err)
				http.Error(w, "Failed to remove association", http.StatusInternalServerError)
				return
			}
		}

		if err := db.DB.Model(&group).Update("Admins", newAdmins).Error; err != nil {
			fmt.Println("Error updating admins list", err)
			http.Error(w, "Failed to update group admins", http.StatusInternalServerError)
			return
		}
	}

	if err := db.DB.Unscoped().Model(&user).Association("Session").Unscoped().Clear(); err != nil {
		fmt.Println("Error deleting session", err)
		http.Error(w, "Failed to delete user session", http.StatusInternalServerError)
		return
	}

	if err := db.DB.Delete(&user).Error; err != nil {
		fmt.Println("Error deleting user", err)
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}
