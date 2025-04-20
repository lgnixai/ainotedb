package service

import (
	"log"
	"net"
	"sync"

	"github.com/gorilla/websocket"
)

// Client represents a single WebSocket client.
type Client struct {
	conn *websocket.Conn
	// TODO: Add subscription information (e.g., subscribed tables/records)
}

// RealtimeService manages WebSocket connections and broadcasts messages.
type RealtimeService struct {
	clients    map[*Client]bool
	broadcast  chan []byte  // Channel for broadcasting messages to all clients
	register   chan *Client // Channel for registering new clients
	unregister chan *Client // Channel for unregistering clients
	mutex      sync.Mutex
}

// NewRealtimeService creates a new RealtimeService.
func NewRealtimeService() *RealtimeService {
	return &RealtimeService{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

// Run starts the RealtimeService message handling loop.
func (s *RealtimeService) Run() {
	for {
		select {
		case client := <-s.register:
			s.mutex.Lock()
			s.clients[client] = true
			s.mutex.Unlock()
			log.Println("Client registered")
		case client := <-s.unregister:
			s.mutex.Lock()
			if _, ok := s.clients[client]; ok {
				delete(s.clients, client)
				close(client.conn.UnderlyingConn().(*net.TCPConn).CloseWrite())
				log.Println("Client unregistered")
			}
			s.mutex.Unlock()
		case message := <-s.broadcast:
			s.mutex.Lock()
			log.Printf("Broadcasting message: %s", string(message))
			for client := range s.clients {
				// TODO: Implement filtering based on client subscriptions
				if err := client.conn.WriteMessage(websocket.TextMessage, message); err != nil {
					log.Printf("Error broadcasting to client: %v", err)
					// Consider unregistering the client if write fails
					// s.unregister <- client
				}
			}
			s.mutex.Unlock()
		}
	}
}

// RegisterClient adds a new client to the service.
func (s *RealtimeService) RegisterClient(conn *websocket.Conn) {
	client := &Client{conn: conn}
	s.register <- client
	// TODO: Start a goroutine to read messages from this client
}

// UnregisterClient removes a client from the service.
func (s *RealtimeService) UnregisterClient(client *Client) {
	s.unregister <- client
}

// BroadcastMessage sends a message to all connected clients.
func (s *RealtimeService) BroadcastMessage(message []byte) {
	s.broadcast <- message
}

// TODO: Add methods for handling specific events (e.g., record updates)
// These methods would likely be called by other services (e.g., RecordService)
// and would then call BroadcastMessage with the relevant data.

// Example: HandleRecordUpdate(event RecordUpdateEvent)
// func (s *RealtimeService) HandleRecordUpdate(event RecordUpdateEvent) {
// 	 // Serialize event data to JSON or other format
// 	 message, err := json.Marshal(event)
// 	 if err != nil {
// 	 	 log.Printf("Error marshalling record update event: %v", err)
// 	 	 return
// 	 }
// 	 s.BroadcastMessage(message)
// }
