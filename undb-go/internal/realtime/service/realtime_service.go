package service

import (
	"encoding/json"
	"log"
	"net"
	"sync"
	"time"

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
				if tcpConn, ok := client.conn.UnderlyingConn().(*net.TCPConn); ok {
    _ = tcpConn.CloseWrite()
}
_ = client.conn.Close()
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


type Operation struct {
	Type      string          `json:"type"`      // insert, update, delete
	TableID   string          `json:"tableId"`
	RecordID  string          `json:"recordId"`
	UserID    string          `json:"userId"`
	Timestamp time.Time       `json:"timestamp"`
	Data      map[string]any  `json:"data"`
}

func (s *RealtimeService) BroadcastOperation(op Operation) {
	message, err := json.Marshal(op)
	if err != nil {
		log.Printf("Failed to marshal operation: %v", err)
		return
	}

	// 记录操作历史
	s.mutex.Lock()
	// TODO: 实现操作历史存储
	s.mutex.Unlock()

	// 广播给所有客户端
	s.broadcast <- message
}

func (s *RealtimeService) HandleConflict(op1, op2 Operation) Operation {
	// 使用OT算法处理冲突
	if op1.Timestamp.Before(op2.Timestamp) {
		return op2
	}
	return op1
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