package realtime

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/gin-gonic/gin"
)

type Event struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type Connection struct {
	conn      *websocket.Conn
	send      chan []byte
	spaceID   string
	userID    string
	lastPing  time.Time
	isAlive   bool
}

type WebSocketManager struct {
	connections map[string][]*Connection
	mutex       sync.RWMutex
}

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		connections: make(map[string][]*Connection),
	}
}

func (m *WebSocketManager) AddConnection(spaceID string, userID string, conn *websocket.Conn) *Connection {
	connection := &Connection{
		conn:     conn,
		send:     make(chan []byte, 256),
		spaceID:  spaceID,
		userID:   userID,
		lastPing: time.Now(),
		isAlive:  true,
	}

	m.mutex.Lock()
	m.connections[spaceID] = append(m.connections[spaceID], connection)
	m.mutex.Unlock()

	// Setup connection
	conn.SetReadLimit(maxMessageSize)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(pongWait))
		connection.lastPing = time.Now()
		return nil
	})

	// Start goroutines for reading and writing
	go connection.writePump()
	go connection.readPump(m)

	return connection
}

func (c *Connection) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			w.Close()

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Connection) readPump(m *WebSocketManager) {
	defer func() {
		m.removeConnection(c)
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				// Log error
			}
			break
		}

		var event Event
		if err := json.Unmarshal(message, &event); err != nil {
			continue
		}

		// Handle different event types
		switch event.Type {
		case "ping":
			c.lastPing = time.Now()
		default:
			// Broadcast to other users in space
			m.BroadcastToOthers(c.spaceID, c.userID, message)
		}
	}
}

func (m *WebSocketManager) removeConnection(conn *Connection) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	connections := m.connections[conn.spaceID]
	for i, c := range connections {
		if c == conn {
			m.connections[conn.spaceID] = append(connections[:i], connections[i+1:]...)
			close(c.send)
			break
		}
	}
}

func (m *WebSocketManager) Broadcast(spaceID string, message []byte) {
	m.mutex.RLock()
	connections := m.connections[spaceID]
	m.mutex.RUnlock()

	for _, conn := range connections {
		if conn.isAlive {
			select {
			case conn.send <- message:
			default:
				m.removeConnection(conn)
			}
		}
	}
}

func (m *WebSocketManager) BroadcastToOthers(spaceID string, senderID string, message []byte) {
	m.mutex.RLock()
	connections := m.connections[spaceID]
	m.mutex.RUnlock()

	for _, conn := range connections {
		if conn.userID != senderID && conn.isAlive {
			select {
			case conn.send <- message:
			default:
				m.removeConnection(conn)
			}
		}
	}
}