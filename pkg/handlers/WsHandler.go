package handlers

import (
	"fmt"
	"net/http"

	"github.com/FKuiv/LocalChat/pkg/controller"
	"github.com/FKuiv/LocalChat/pkg/utils"
	"github.com/FKuiv/LocalChat/pkg/websocket"
	gorillaWs "github.com/gorilla/websocket"
)

type wsHandler struct {
	UserController    controller.UserController
	GroupController   controller.GroupController
	MessageController controller.MessageController
}

func NewWsHandler(controllers controller.Controllers) *wsHandler {
	return &wsHandler{
		UserController:    *controllers.UserController,
		GroupController:   *controllers.GroupController,
		MessageController: *controllers.MessageController,
	}
}

var wsConnUpgrader = gorillaWs.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // Currently let everyone to connect
}

func (handler *wsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	userCookie, cookieErr := utils.GetUserCookie(r)
	if utils.CookieError(cookieErr, w) {
		return
	}

	user, userErr := handler.UserController.Service.GetUserById(userCookie.Value)

	if userErr != nil {
		http.Error(w, fmt.Sprintf("Error getting user: %s", userErr), http.StatusInternalServerError)
		return
	}

	conn, err := wsConnUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	groupIds, groupsErr := handler.GroupController.Service.GetAllUserGroupIds(user.ID)
	if groupsErr != nil {
		http.Error(w, fmt.Sprintf("%s", groupsErr), http.StatusInternalServerError)
		return
	}

	client := &websocket.Client{GroupIds: groupIds, Hub: handler.UserController.Service.GetWsHub(), Socket: conn, Send: make(chan websocket.WsMessage), User: *user}
	client.Hub.Register <- client

	go client.Write()
	go client.Read()
}

func (handler *wsHandler) RefreshWs(w http.ResponseWriter, r *http.Request) {

}
