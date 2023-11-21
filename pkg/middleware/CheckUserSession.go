package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/FKuiv/LocalChat/pkg/handlers"
	"github.com/FKuiv/LocalChat/pkg/models"
	"gorm.io/gorm"
)

func CheckUserSession(next http.Handler, db handlers.DBHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip these endpoints
		if r.RequestURI == "/login" || r.RequestURI == "/user" {
			next.ServeHTTP(w, r)
			return
		}

		var session models.Session
		userId := r.Header.Get("UserId")
		result := db.DB.First(&session, "user_id = ?", userId)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, fmt.Sprintf("User with ID: %s does not have a session", userId), http.StatusUnauthorized)
			return
		}

		switch comparison := session.ExpiresAt.Compare(time.Now()); comparison {
		case -1:
			// WRONG. The session has expired
			http.Error(w, "User session expired", http.StatusUnauthorized)
			db.DB.Delete(&session, session.ID)
			return
		default:
			// The session is valid
			next.ServeHTTP(w, r)
		}
	})
}
