package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/FKuiv/LocalChat/pkg/models"
	"github.com/FKuiv/LocalChat/pkg/utils"
	"gorm.io/gorm"
)

func CheckUserSession(next http.Handler, db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip these endpoints
		if r.RequestURI == "/login" || r.RequestURI == "/user" {
			next.ServeHTTP(w, r)
			return
		}

		var session models.Session

		cookies, err := utils.GetCookies(r)
		if utils.CookieError(err, w) {
			return
		}

		result := db.First(&session, "id = ?", cookies.Session.Value)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, fmt.Sprintf("User with ID: %s does not have a session", cookies.User.Value), http.StatusUnauthorized)
			return
		}

		if session.UserID != cookies.User.Value {
			http.Error(w, "User ID and session ID do not match", http.StatusForbidden)
			return
		}

		switch comparison := session.ExpiresAt.Compare(time.Now()); comparison {
		case -1:
			// WRONG. The session has expired
			http.Error(w, "User session expired", http.StatusUnauthorized)
			db.Delete(&session, session.ID)
			return
		default:
			// The session is valid
			next.ServeHTTP(w, r)
		}
	})
}
