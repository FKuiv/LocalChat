package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/FKuiv/LocalChat/pkg/controller"
	"github.com/FKuiv/LocalChat/pkg/utils"
)

func CheckUserSession(next http.Handler, controllers *controller.Controllers) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip these endpoints
		if r.RequestURI == "/login" || r.RequestURI == "/user" && r.Method == http.MethodPost {
			next.ServeHTTP(w, r)
			return
		}

		cookies, err := utils.GetCookies(r)
		if utils.CookieError(err, w) {
			return
		}

		session, sessionErr := controllers.UserController.Service.GetSessionById(cookies.Session.Value, cookies.User.Value)

		if strings.Contains(fmt.Sprintf("%s", sessionErr), "User with ID") {
			utils.DeleteCookies(w)
			http.Error(w, fmt.Sprintf("%s", sessionErr), http.StatusUnauthorized)
			return
		}

		if session.UserID != cookies.User.Value {
			utils.DeleteCookies(w)
			http.Error(w, "User ID and session ID do not match", http.StatusForbidden)
			return
		}

		if sessionErr != nil {
			utils.DeleteCookies(w)
			http.Error(w, fmt.Sprintf("Failed to get session: %s", sessionErr), http.StatusInternalServerError)
			return
		}

		switch comparison := session.ExpiresAt.Compare(time.Now()); comparison {
		case -1:
			// WRONG. The session has expired
			deleteErr := controllers.UserController.Service.DeleteSession(cookies.Session.Value, cookies.User.Value)
			if deleteErr != nil {
				http.Error(w, fmt.Sprintf("Failed to delete session: %s", deleteErr), http.StatusInternalServerError)
				return
			}
			utils.DeleteCookies(w)
			http.Error(w, "User session expired", http.StatusUnauthorized)
			return
		default:
			// The session is valid
			next.ServeHTTP(w, r)
		}
	})
}
