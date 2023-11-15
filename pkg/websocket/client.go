package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// Representation of the connection to the end user
type Client struct {
	socket *websocket.Conn

	// Channel for sending and receving messages from other clients
	send   chan []byte
	groups []*Group
}

func (c *Client) ID() string {
	return fmt.Sprintf("%p", c)
}

func (c *Client) read() {
	defer c.socket.Close()

	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		c.socket.WriteJSON(message)
	}
}
