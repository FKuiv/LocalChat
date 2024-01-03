package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/FKuiv/LocalChat/pkg/controller"
	"github.com/FKuiv/LocalChat/pkg/models"
	"github.com/FKuiv/LocalChat/pkg/utils"
	"github.com/gorilla/mux"
)

type userHandler struct {
	UserController controller.UserController
}

func NewUserHandler(controller controller.UserController) *userHandler {
	return &userHandler{
		UserController: controller,
	}
}

func (handler *userHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := handler.UserController.Service.GetAllUsers()

	if err != nil {
		http.Error(w, fmt.Sprintf("There was an error getting users: %s", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func (handler *userHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, idOk := vars["id"]

	if utils.MuxVarsNotProvided(idOk, userId, "User ID", w) {
		return
	}

	user, err := handler.UserController.Service.GetUserById(userId)

	if err != nil {
		http.Error(w, "Error getting user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (handler *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userInfo models.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userInfo)
	if utils.DecodingErr(err, "/user", w) {
		return
	}

	newUser, err := handler.UserController.Service.CreateUser(userInfo)
	errString := fmt.Sprintf("%s", err)

	if err != nil && strings.Contains(errString, "already taken") {
		http.Error(w, errString, http.StatusConflict)
		return
	}

	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating user: %s", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newUser)
}

func (handler *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userCookie, cookieErr := utils.GetUserCookie(r)
	if utils.CookieError(cookieErr, w) {
		return
	}

	err := handler.UserController.Service.DeleteUser(userCookie.Value)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting user: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}

func (handler *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var newUserInfo models.UserRequest
	err := json.NewDecoder(r.Body).Decode(&newUserInfo)
	if utils.DecodingErr(err, "/user", w) {
		return
	}
	userCookie, cookieErr := utils.GetUserCookie(r)
	if utils.CookieError(cookieErr, w) {
		return
	}

	currentUser, err := handler.UserController.Service.UpdateUser(newUserInfo, userCookie.Value)

	if err != nil && strings.Contains(fmt.Sprintf("%s", err), "Username already exists") {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating user: %s", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(currentUser)
}

func (handler *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var userInfo models.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userInfo)
	if utils.DecodingErr(err, "/login", w) {
		return
	}

	session, err := handler.UserController.Service.CreateSession(userInfo)

	if err != nil && strings.Contains(fmt.Sprintf("%s", err), "Wrong password") {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusUnauthorized)
		return
	}

	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating session: %s", err), http.StatusInternalServerError)
		return
	}

	// 604800 is 7 days in seconds. Using MaxAge because Safari prefers it. Just in case setting expires as well
	sessionCookie := http.Cookie{Name: "Session", Value: session.ID, Domain: "localhost", Path: "/", Expires: session.ExpiresAt, MaxAge: 604800, HttpOnly: true}
	http.SetCookie(w, &sessionCookie)

	userCookie := http.Cookie{Name: "UserId", Value: session.UserID, Domain: "localhost", Path: "/", Expires: session.ExpiresAt, MaxAge: 604800, HttpOnly: false}
	http.SetCookie(w, &userCookie)
}

func (handler *userHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookies, cookiesErr := utils.GetCookies(r)

	if utils.CookieError(cookiesErr, w) {
		return
	}

	err := handler.UserController.Service.DeleteSession(cookies.Session.Value, cookies.User.Value)

	if err != nil && strings.Contains(fmt.Sprintf("%s", err), "Forbidden") {
		http.Error(w, "User does not own this session", http.StatusForbidden)
		return
	}

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete session: %s", err), http.StatusInternalServerError)
		return
	}

	// Deleting cookies
	utils.DeleteCookies(w)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Session deleted successfully"))
}

func (handler *userHandler) UploadProfilePic(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(utils.MULTIPART_FORM_MAX_MEMORY); err != nil {
		http.Error(w, fmt.Sprintf("error parsing request: %s", err), http.StatusBadRequest)
		return
	}
	userCookie, cookieErr := utils.GetUserCookie(r)
	if utils.CookieError(cookieErr, w) {
		return
	}

	// Limit upload size
	r.Body = http.MaxBytesReader(w, r.Body, utils.MULTIPART_FORM_MAX_MEMORY) // 5 Mb

	file, multipartFileHeader, err := r.FormFile("picture")
	if err != nil {
		http.Error(w, fmt.Sprintf("error parsing request: %s", err), http.StatusBadRequest)
		return
	}

	if err := handler.UserController.Service.SaveProfilePic(file, multipartFileHeader, userCookie.Value); err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File uploaded successfully"))
}

func (handler *userHandler) GetProfilePic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, idOk := vars["id"]

	if utils.MuxVarsNotProvided(idOk, userId, "User ID", w) {
		return
	}

	picUrl, err := handler.UserController.Service.GetProfilePic(userId)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(picUrl))
}

func (handler *userHandler) GetUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, idOk := vars["id"]

	if utils.MuxVarsNotProvided(idOk, userId, "User ID", w) {
		return
	}

	username, err := handler.UserController.Service.GetUsername(userId)

	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(username))
}
