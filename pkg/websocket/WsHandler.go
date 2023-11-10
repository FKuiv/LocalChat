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

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsConnUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("Received message:", string(message))

		err = conn.WriteMessage(messageType, []byte("Hi from backend"))

		if err != nil {
			log.Println(err)
			return
		}
	}
}
