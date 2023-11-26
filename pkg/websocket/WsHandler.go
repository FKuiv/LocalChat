package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var wsConnUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // Currently let everyone to connect
}

func WsHandler(hub *Hub, w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("UserId")
	conn, err := wsConnUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		fmt.Println(err)
		return
	}

	client := &Client{hub: hub, ID: userId, socket: conn, send: make(chan []byte, 256)}
	fmt.Println("New ws client:", client)
	client.hub.register <- client

	go client.write()
	go client.read()
}
