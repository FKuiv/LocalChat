package handlers

import (
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

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}
