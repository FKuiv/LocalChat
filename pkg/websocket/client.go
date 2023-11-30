package websocket

import (
	"fmt"
	"log"

	"github.com/FKuiv/LocalChat/pkg/models"
	"github.com/gorilla/websocket"
)

// Representation of the connection to the end user
type Client struct {
	models.User
	socket *websocket.Conn
	hub    *Hub
	send   chan []byte
}

// This specific user is sending a message
func (c *Client) read() {
	defer func() {
		c.hub.unregister <- c
		c.socket.Close()
	}()

	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			log.Println(err)
			fmt.Println(err)
			return
		}

		fmt.Println("Reading message:", []byte(message), "with user:", c.Username)
		c.hub.broadcast <- message
	}
}

// A new message is broadcasted to every user
func (c *Client) write() {
	defer c.socket.Close()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			fmt.Println("This message:", []byte(message), "for user:", c.Username)
			err := c.socket.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				fmt.Printf("Error writing to WebSocket for client %s: %v\n", c.ID, err)
				return
			}
		}
	}
}
