package websocket

import (
	"log"

	"github.com/FKuiv/LocalChat/pkg/models"
	"github.com/gorilla/websocket"
)

// Representation of the connection to the end user
type Client struct {
	models.User
	Socket   *websocket.Conn
	Hub      *Hub
	Send     chan models.MessageRequest
	GroupIds []string
}

// This specific user is sending a message
func (c *Client) Read() {
	defer func() {
		c.Hub.Unregister <- c
		c.Socket.Close()
	}()

	for {

		var message models.MessageRequest
		err := c.Socket.ReadJSON(&message)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("Reading message:", message, "with user:", c.Username)
		c.Hub.Broadcast <- message

		_, dbErr := c.Hub.controllers.MessageController.Service.CreateMessage(models.MessageRequest{GroupID: message.GroupID, Content: message.Content}, c.ID)
		if dbErr != nil {
			log.Println("Error saving message:", dbErr)
		}
	}
}

// A new message is broadcasted to every user
func (c *Client) Write() {
	defer c.Socket.Close()

	for message := range c.Send {
		err := c.Socket.WriteJSON(&message)
		if err != nil {
			log.Printf("Error writing to WebSocket for client %s: %v\n", c.ID, err)
			return
		}
	}
}
