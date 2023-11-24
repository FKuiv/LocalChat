package handlers

import (
	"fmt"
	"net/http"

	"github.com/FKuiv/LocalChat/pkg/models"
	"github.com/FKuiv/LocalChat/pkg/utils"
	"github.com/gorilla/mux"
)

func (db DBHandler) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	groupId, idOk := mux.Vars(r)["id"]

	if utils.MuxVarsNotProvided(idOk, groupId, "Group ID", w) {
		return
	}

	userId := r.Header.Get("UserId")

	var group models.Group
	result := db.DB.Where("id = ?", groupId).First(&group)

	if utils.ItemNotFound(result.Error, "Group", w) {
		return
	}

	isAdmin := false
	for _, adminId := range group.Admins {
		if userId == adminId {
			isAdmin = true
			break
		}
	}

	if !isAdmin {
		http.Error(w, "User needs to be admin to delete this group", http.StatusForbidden)
		return
	}

	// Need to delete all the messages inside a group first
	if err := db.DB.Unscoped().Model(&group).Association("Messages").Unscoped().Clear(); err != nil {
		fmt.Println("Error deleting all the messages in a group", err)
		http.Error(w, "Failed to delete all group messages", http.StatusInternalServerError)
		return
	}

	if err := db.DB.Select("Users").Delete(&group).Error; err != nil {
		fmt.Println("Error removing references from user_groups table", err)
		http.Error(w, "Failed to remove references from user_groups table", http.StatusInternalServerError)
		return
	}

	// Remove references to the group from the user_groups join table
	// if err := db.DB.Model(&models.User{}).Association("Groups").Delete(&group); err != nil {
	// 	fmt.Println("Error removing references from user_groups table", err)
	// 	http.Error(w, "Failed to remove references from user_groups table", http.StatusInternalServerError)
	// 	return
	// }

	// if err := db.DB.Delete(&group).Error; err != nil {
	// 	fmt.Println("Error deleting group:", err)
	// 	http.Error(w, "Failed to delete group", http.StatusInternalServerError)
	// 	return
	// }

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Group deleted successfully"))
}
