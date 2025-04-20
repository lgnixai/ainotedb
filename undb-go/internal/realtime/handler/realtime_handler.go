package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/undb/undb-go/internal/realtime/service"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections by default
		// TODO: Implement proper origin checking for security
		return true
	},
}

// RealtimeHandler handles WebSocket connections.
type RealtimeHandler struct {
	svc *service.RealtimeService
}

// NewRealtimeHandler creates a new RealtimeHandler.
func NewRealtimeHandler(svc *service.RealtimeService) *RealtimeHandler {
	return &RealtimeHandler{svc: svc}
}

// ServeWS handles WebSocket requests.
func (h *RealtimeHandler) ServeWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}
	defer conn.Close()

	log.Println("WebSocket client connected")

	// TODO: Implement message handling loop (read messages, manage subscriptions)
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("WebSocket read error: %v", err)
			break
		}
		log.Printf("Received message type %d: %s", messageType, string(p))

		// Echo message back for now
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Printf("WebSocket write error: %v", err)
			break
		}
	}

	log.Println("WebSocket client disconnected")
}
