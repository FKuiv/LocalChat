package utils

import "net/http"

func DeleteCookies(w http.ResponseWriter) {
	// Deleting cookies
	sessionCookie := http.Cookie{Name: "Session", Value: "", Domain: "localhost", Path: "/", MaxAge: -1, HttpOnly: true}
	http.SetCookie(w, &sessionCookie)

	userCookie := http.Cookie{Name: "UserId", Value: "", Domain: "localhost", Path: "/", MaxAge: -1, HttpOnly: true}
	http.SetCookie(w, &userCookie)
}
