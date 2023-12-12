package websocket

import (
	"log"
	"sync"

	"github.com/FKuiv/LocalChat/pkg/controller"
)

type Hub struct {
	Clients     map[string]*Client
	Groups      map[string]*WsGroup
	Register    chan *Client
	Unregister  chan *Client
	Broadcast   chan WsMessage
	Refresh     chan RefreshMessage
	controllers *controller.Controllers
	mutex       sync.Mutex
}

type WsGroup struct {
	ID      string
	Clients map[string]*Client
}

type WsMessage struct {
	UserID  string `json:"user_id"`  // aka the Author of the message
	GroupID string `json:"group_id"` // to filter the message into the right chat
	Content string `json:"content"`
}

type RefreshMessage struct {
	NewGroupId      string   `json:"new_group_id"`
	ClientsToUpdate []string `json:"clients_to_update"`
}

func NewHub(controllers *controller.Controllers) *Hub {
	return &Hub{
		Clients:     make(map[string]*Client),
		Groups:      make(map[string]*WsGroup),
		Register:    make(chan *Client),
		Unregister:  make(chan *Client),
		Broadcast:   make(chan WsMessage),
		Refresh:     make(chan RefreshMessage),
		controllers: controllers,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.mutex.Lock()

			h.Clients[client.ID] = client

			for _, groupId := range client.GroupIds {
				group, exists := h.Groups[groupId]
				if !exists {
					group = &WsGroup{ID: groupId, Clients: make(map[string]*Client)}
					h.Groups[groupId] = group
					log.Printf("\nGroup %s created. Total groups: %d", group.ID, len(h.Groups))
				}

				log.Printf("\nUser %s registered to group %s. Total groups: %d", client.ID, group.ID, len(h.Groups))
				group.Clients[client.ID] = client
			}

			h.mutex.Unlock()
			log.Printf("\nUser %s registered. Total users: %d", client.Username, len(h.Clients))

		case client := <-h.Unregister:
			h.mutex.Lock()
			delete(h.Clients, client.ID)
			for _, groupId := range client.GroupIds {
				group := h.Groups[groupId]
				delete(group.Clients, client.ID)
			}
			close(client.Send)
			h.mutex.Unlock()
			log.Printf("\nUser %s unregistered. Total users: %d", client.Username, len(h.Clients))

		case message := <-h.Broadcast:
			h.mutex.Lock()
			if group, exists := h.Groups[message.GroupID]; exists {
				for _, client := range group.Clients {
					select {
					case client.Send <- message:
					default:
						close(client.Send)
						delete(h.Clients, client.ID)
					}
				}
			}
			h.mutex.Unlock()

		case message := <-h.Refresh:
			h.mutex.Lock()
			group, exists := h.Groups[message.NewGroupId]

			if !exists {
				group = &WsGroup{ID: message.NewGroupId, Clients: make(map[string]*Client)}
				h.Groups[message.NewGroupId] = group
				log.Printf("\nGroup %s created. Total groups: %d", group.ID, len(h.Groups))
			}

			for _, clientId := range message.ClientsToUpdate {
				client := h.Clients[clientId]
				group.Clients[clientId] = client
			}

			h.mutex.Unlock()
		}

	}
}
