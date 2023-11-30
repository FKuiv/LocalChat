package websocket

import (
	"fmt"
	"net/http"

	"github.com/FKuiv/LocalChat/pkg/controller"
	"github.com/gorilla/websocket"
)

var wsConnUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // Currently let everyone to connect
}

func WsHandler(hub *Hub, controllers *controller.Controllers, w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("UserId")
	user, userErr := controllers.UserController.Service.GetUserById(userId)

	if userErr != nil {
		http.Error(w, fmt.Sprintf("Error getting user: %s", userErr), http.StatusInternalServerError)
		return
	}

	conn, err := wsConnUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &Client{hub: hub, socket: conn, send: make(chan []byte, 256), User: *user}
	client.hub.register <- client

	go client.write()
	go client.read()
}
