package router

import (
	"github.com/gin-gonic/gin"

	"github.com/undb/undb-go/internal/realtime/handler"
)

// RegisterRealtimeRoutes registers the WebSocket route.
func RegisterRoutes(r *gin.Engine, h *handler.RealtimeHandler) {
	r.GET("/ws", func(c *gin.Context) {
		h.ServeWS(c.Writer, c.Request)
	})
}
