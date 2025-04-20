package realtime

import (
	"sync"
	"github.com/gorilla/websocket"
	"github.com/gin-gonic/gin"
)

type Connection struct {
	conn *websocket.Conn
	send chan []byte
}

type WebSocketManager struct {
	connections map[string][]*Connection
	mutex       sync.RWMutex
}

func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		connections: make(map[string][]*Connection),
	}
}

func (m *WebSocketManager) AddConnection(spaceID string, conn *websocket.Conn) *Connection {
	connection := &Connection{
		conn: conn,
		send: make(chan []byte, 256),
	}
	
	m.mutex.Lock()
	m.connections[spaceID] = append(m.connections[spaceID], connection)
	m.mutex.Unlock()
	
	return connection
}

func (m *WebSocketManager) Broadcast(spaceID string, message []byte) {
	m.mutex.RLock()
	connections := m.connections[spaceID]
	m.mutex.RUnlock()
	
	for _, conn := range connections {
		select {
		case conn.send <- message:
		default:
			close(conn.send)
		}
	}
}