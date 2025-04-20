package router

import (
	"github.com/gin-gonic/gin"
	"github.com/undb/undb-go/internal/table/handler"
)

// RegisterTableRoutes 注册表格路由
func RegisterTableRoutes(r *gin.RouterGroup, h *handler.TableHandler) {
	tables := r.Group("/tables")
	{
		tables.POST("", h.CreateTable)
		tables.GET("/:id", h.GetTable)
		tables.GET("/space/:space_id", h.GetTables)
		tables.PUT("/:id", h.UpdateTable)
		tables.DELETE("/:id", h.DeleteTable)
	}
}
