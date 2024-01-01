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

type groupHandler struct {
	GroupController controller.GroupController
}

func NewGroupHandler(controller controller.GroupController) *groupHandler {
	return &groupHandler{
		GroupController: controller,
	}
}

func (handler *groupHandler) GetAllGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := handler.GroupController.Service.GetAllGroups()

	if err != nil {
		http.Error(w, fmt.Sprintf("There was an error getting groups: %s", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(groups)
}

func (handler *groupHandler) GetGroupById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupId, idOk := vars["id"]

	if utils.MuxVarsNotProvided(idOk, groupId, "Group ID", w) {
		return
	}

	group, err := handler.GroupController.Service.GetGroupById(groupId)

	if err != nil {
		http.Error(w, "Error getting group", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(group)
}

func (handler *groupHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	var group models.GroupRequest
	err := json.NewDecoder(r.Body).Decode(&group)
	if utils.DecodingErr(err, "/group", w) {
		return
	}
	fmt.Println(group)
	newGroup, err := handler.GroupController.Service.CreateGroup(group)

	errString := fmt.Sprintf("%s", err)

	if err != nil && (strings.Contains(errString, "Group name can't be empty") || strings.Contains(errString, "There needs to be at least 1 admin and user in group")) {
		http.Error(w, errString, http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating group: %s", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newGroup)
}

func (handler *groupHandler) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	groupId, idOk := mux.Vars(r)["id"]

	if utils.MuxVarsNotProvided(idOk, groupId, "Group ID", w) {
		return
	}

	userCookie, cookieErr := utils.GetUserCookie(r)
	if utils.CookieError(cookieErr, w) {
		return
	}

	err := handler.GroupController.Service.DeleteGroup(groupId, userCookie.Value)

	if err != nil && strings.Contains(fmt.Sprintf("%s", err), "User needs to be admin to delete this group") {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusForbidden)
		return
	}

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete group: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Group deleted successfully"))
}

func (handler *groupHandler) UpdateGroup(w http.ResponseWriter, r *http.Request) {
	var newGroupInfo models.GroupRequest
	err := json.NewDecoder(r.Body).Decode(&newGroupInfo)
	if utils.DecodingErr(err, "/group", w) {
		return
	}

	vars := mux.Vars(r)
	groupId, idOk := vars["id"]

	if utils.MuxVarsNotProvided(idOk, groupId, "Group ID", w) {
		return
	}

	currentGroup, err := handler.GroupController.Service.UpdateGroup(newGroupInfo, groupId)

	errString := fmt.Sprintf("%s", err)
	if err != nil && strings.Contains(errString, "not found") {
		http.Error(w, errString, http.StatusNotFound)
		return
	}

	if err != nil && strings.Contains(errString, "cannot") {
		http.Error(w, errString, http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, errString, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(currentGroup)
}

func (handler *groupHandler) GetAllUserGroups(w http.ResponseWriter, r *http.Request) {
	userCookie, cookieErr := utils.GetUserCookie(r)

	if utils.CookieError(cookieErr, w) {
		return
	}

	userGroups, err := handler.GroupController.Service.GetAllUserGroups(userCookie.Value)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching user groups: %s", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(userGroups)
}

func (handler *groupHandler) UploadGroupPic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupId, idOk := vars["id"]

	if utils.MuxVarsNotProvided(idOk, groupId, "Group ID", w) {
		return
	}

	if err := r.ParseMultipartForm(utils.MULTIPART_FORM_MAX_MEMORY); err != nil {
		http.Error(w, fmt.Sprintf("error parsing request: %s", err), http.StatusBadRequest)
		return
	}

	// Limit upload size
	r.Body = http.MaxBytesReader(w, r.Body, utils.MULTIPART_FORM_MAX_MEMORY) // 5 Mb

	file, multipartFileHeader, err := r.FormFile("picture")
	if err != nil {
		http.Error(w, fmt.Sprintf("error parsing request: %s", err), http.StatusBadRequest)
		return
	}

	if err := handler.GroupController.Service.SaveGroupPic(file, multipartFileHeader, groupId); err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File uploaded successfully"))
}

func (handler *groupHandler) GetGroupPic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupId, idOk := vars["id"]

	if utils.MuxVarsNotProvided(idOk, groupId, "Group ID", w) {
		return
	}

	picUrl, err := handler.GroupController.Service.GetGroupPic(groupId)

	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(picUrl))
}
