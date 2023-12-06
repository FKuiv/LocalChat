package utils

import (
	"fmt"
	"net/http"
)

type LocalChatCookies struct {
	User    *http.Cookie
	Session *http.Cookie
}

func GetCookies(r *http.Request) (*LocalChatCookies, error) {
	sessionCookie, sessionCookieErr := r.Cookie("Session")
	if sessionCookieErr != nil {
		return nil, &CustomError{Message: fmt.Sprintf("Error with session cookie: %s", sessionCookieErr)}
	}

	userCookie, userCookieErr := r.Cookie("UserId")
	if userCookieErr != nil {
		return nil, &CustomError{Message: fmt.Sprintf("Error with user cookie: %s", userCookieErr)}
	}

	return &LocalChatCookies{User: userCookie, Session: sessionCookie}, nil
}

func GetUserCookie(r *http.Request) (*http.Cookie, error) {
	userCookie, userCookieErr := r.Cookie("UserId")
	if userCookieErr != nil {
		return nil, &CustomError{Message: fmt.Sprintf("Error with user cookie: %s", userCookieErr)}
	}

	return userCookie, nil
}

func GetSessionCookie(r *http.Request) (*http.Cookie, error) {
	sessionCookie, sessionCookieErr := r.Cookie("Session")
	if sessionCookieErr != nil {
		return nil, &CustomError{Message: fmt.Sprintf("Error with session cookie: %s", sessionCookieErr)}
	}

	return sessionCookie, nil
}
