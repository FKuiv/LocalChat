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

type messageHandler struct {
	MessageController controller.MessageController
}

func NewMessageHandler(controller controller.MessageController) *messageHandler {
	return &messageHandler{
		MessageController: controller,
	}
}

func (handler *messageHandler) GetAllMessages(w http.ResponseWriter, r *http.Request) {
	messages, err := handler.MessageController.Service.GetAllMessages()

	if err != nil {
		http.Error(w, fmt.Sprintf("There was an error getting messages: %s", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(messages)
}

func (handler *messageHandler) GetMessageById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	messageId, idOk := vars["id"]

	if utils.MuxVarsNotProvided(idOk, messageId, "Message ID", w) {
		return
	}

	message, err := handler.MessageController.Service.GetMessageById(messageId)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting group: %s", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(message)
}

func (handler *messageHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var message models.MessageRequest
	err := json.NewDecoder(r.Body).Decode(&message)
	if utils.DecodingErr(err, "/message", w) {
		return
	}

	userCookie, cookieErr := utils.GetUserCookie(r)
	if utils.CookieError(cookieErr, w) {
		return
	}

	newMessage, err := handler.MessageController.Service.CreateMessage(message, userCookie.Value)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating message: %s", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newMessage)
}

func (handler *messageHandler) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	messageId, idOk := mux.Vars(r)["id"]
	if utils.MuxVarsNotProvided(idOk, messageId, "Message ID", w) {
		return
	}

	userCookie, cookieErr := utils.GetUserCookie(r)
	if utils.CookieError(cookieErr, w) {
		return
	}

	err := handler.MessageController.Service.DeleteMessage(messageId, userCookie.Value)

	if err != nil && strings.Contains(fmt.Sprintf("%s", err), "does not own this message") {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusForbidden)
		return
	}

	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message deleted successfully"))
}

func (handler *messageHandler) UpdateMessage(w http.ResponseWriter, r *http.Request) {
	var newMessageInfo models.UpdateMessage
	err := json.NewDecoder(r.Body).Decode(&newMessageInfo)
	if utils.DecodingErr(err, "/message", w) {
		return
	}
	vars := mux.Vars(r)
	messageId, idOk := vars["id"]

	if utils.MuxVarsNotProvided(idOk, messageId, "Message ID", w) {
		return
	}

	updatedMessage, err := handler.MessageController.Service.UpdateMessage(newMessageInfo, messageId)

	errString := fmt.Sprintf("%s", err)
	if err != nil && strings.Contains(errString, "cannot have empty content") {
		http.Error(w, errString, http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, errString, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedMessage)
}
