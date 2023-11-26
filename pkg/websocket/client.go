package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// Representation of the connection to the end user
type Client struct {
	socket *websocket.Conn
	hub    *Hub
	ID     string
	// Channel for sending and receving messages from other clients
	send   chan []byte
	groups []*Group
}

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

		fmt.Println("Reading message:", message, "with user:", c.ID)
		c.hub.broadcast <- message
	}
}

func (c *Client) write() {
	defer c.socket.Close()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			fmt.Println("This message:", message, "for user:", c.ID)
			err := c.socket.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				fmt.Printf("Error writing to WebSocket for client %s: %v\n", c.ID, err)
				return
			}
		}
	}
}
