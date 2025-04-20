package router

import (
	"github.com/gin-gonic/gin"

	"github.com/undb/undb-go/internal/table/handler"
)

// RegisterTableRoutes 注册表格路由
import "github.com/undb/undb-go/internal/user/middleware"

func RegisterRoutes(r *gin.RouterGroup, h *handler.TableHandler) {
	tables := r.Group("/tables")
	tables.Use(middleware.AuthMiddleware())
	{
		tables.POST("", h.CreateTable)
		tables.GET("/:id", h.GetTable)
		tables.GET("/space/:space_id", h.GetTables)
		tables.PUT("/:id", h.UpdateTable)
		tables.DELETE("/:id", h.DeleteTable)
	}
}
