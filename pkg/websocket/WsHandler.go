package websocket

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FKuiv/LocalChat/pkg/utils"
	gorillaWs "github.com/gorilla/websocket"
)

var wsConnUpgrader = gorillaWs.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // Currently let everyone to connect
}

func (hub *Hub) Handle(w http.ResponseWriter, r *http.Request) {
	userCookie, cookieErr := utils.GetUserCookie(r)
	if utils.CookieError(cookieErr, w) {
		return
	}

	user, userErr := hub.controllers.UserController.Service.GetUserById(userCookie.Value)

	if userErr != nil {
		http.Error(w, fmt.Sprintf("Error getting user: %s", userErr), http.StatusInternalServerError)
		return
	}

	conn, err := wsConnUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	groupIds, groupsErr := hub.controllers.GroupController.Service.GetAllUserGroupIds(user.ID)
	if groupsErr != nil {
		http.Error(w, fmt.Sprintf("%s", groupsErr), http.StatusInternalServerError)
		return
	}

	client := &Client{GroupIds: groupIds, Hub: hub, Socket: conn, Send: make(chan WsMessage), User: *user}
	client.Hub.Register <- client

	go client.Write()
	go client.Read()
}

func (hub *Hub) RefreshWs(w http.ResponseWriter, r *http.Request) {
	var message RefreshMessage
	err := json.NewDecoder(r.Body).Decode(&message)
	if utils.DecodingErr(err, "/ws/refresh", w) {
		return
	}

	hub.Refresh <- message

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Websocket refreshed"))
}
