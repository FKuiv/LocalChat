package websocket

import (
	"log"
	"sync"

	"github.com/FKuiv/LocalChat/pkg/models"
)

type Hub struct {
	Clients    map[string]*Client
	Groups     map[string]*WsGroup
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan WsMessage
	mutex      sync.Mutex
}

type WsGroup struct {
	models.Group
	clients map[string]*Client
}

type WsMessage struct {
	UserID  string `json:"user_id"`  // aka the Author of the message
	GroupID string `json:"group_id"` // to filter the message into the right chat
	Content string `json:"content"`
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[string]*Client),
		Groups:     make(map[string]*WsGroup),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan WsMessage),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.mutex.Lock()

			h.Clients[client.ID] = client

			// for groupId := range client.wsgroups {
			// 	group, exists := h.groups[groupId]
			// 	if !exists {
			// 		group = &WsGroup{Group: , clients: make(map[string]*Client)}
			// 		h.groups[groupID] = group
			// 	}
			// 	group.clients[client.ID] = client
			// }

			h.mutex.Unlock()
			log.Printf("User %s registered. Total users: %d", client.Username, len(h.Clients))

		case client := <-h.Unregister:
			h.mutex.Lock()
			delete(h.Clients, client.ID)
			close(client.Send)
			h.mutex.Unlock()
			log.Printf("User %s unregistered. Total users: %d", client.Username, len(h.Clients))

		case message := <-h.Broadcast:
			h.mutex.Lock()
			for _, client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client.ID)
				}
			}
			h.mutex.Unlock()
		}
	}
}
