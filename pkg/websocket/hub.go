package websocket

import (
	"log"
	"sync"
)

type Hub struct {
	clients    map[string]*Client
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
	mutex      sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[string]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mutex.Lock()
			h.clients[client.ID] = client
			h.mutex.Unlock()
			log.Printf("User %s registered. Total users: %d", client.ID, len(h.clients))

		case client := <-h.unregister:
			h.mutex.Lock()
			delete(h.clients, client.ID)
			close(client.send)
			h.mutex.Unlock()
			log.Printf("User %s unregistered. Total users: %d", client.ID, len(h.clients))

		case message := <-h.broadcast:
			h.mutex.Lock()
			for _, client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client.ID)
				}
			}
			h.mutex.Unlock()
		}
	}
}
