package chat

import (
	"fmt"
	"log"
)

// Hub maintains the application state. Clients communicate with each other via the hub
type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

// NewHub creates a new Hub
func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

// Run start the hub listening to channel events
func (h *Hub) Run() {
	fmt.Println("HUB RUNNING")
	for {
		select {
		case client := <-h.Register:
			log.Println("client registering from ", client.Conn.RemoteAddr())
			h.Clients[client] = true
		case client := <-h.Unregister:
			log.Println("client unregistering from ", client.Conn.RemoteAddr())
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			log.Println("broadcasting message: ", message)
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}
