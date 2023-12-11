package websocket

import (
	"fmt"

	"github.com/FKuiv/LocalChat/pkg/models"
	"github.com/gorilla/websocket"
)

// Representation of the connection to the end user
type Client struct {
	models.User
	Socket   *websocket.Conn
	Hub      *Hub
	Send     chan WsMessage
	GroupIds []string
}

// This specific user is sending a message
func (c *Client) Read() {
	defer func() {
		c.Hub.Unregister <- c
		c.Socket.Close()
	}()

	for {

		var message WsMessage
		err := c.Socket.ReadJSON(&message)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Reading message:", message, "with user:", c.Username)
		c.Hub.Broadcast <- message
	}
}

// A new message is broadcasted to every user
func (c *Client) Write() {
	defer c.Socket.Close()

	for message := range c.Send {
		err := c.Socket.WriteJSON(&message)
		if err != nil {
			fmt.Printf("Error writing to WebSocket for client %s: %v\n", c.ID, err)
			return
		}
	}
}
