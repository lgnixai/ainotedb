package router

import (
	"github.com/gin-gonic/gin"

	"github.com/undb/undb-go/internal/field/handler"
)

// RegisterFieldRoutes 注册字段路由
import "github.com/undb/undb-go/internal/user/middleware"

func RegisterRoutes(r *gin.RouterGroup, h *handler.FieldHandler) {
	fields := r.Group("/fields")
	fields.Use(middleware.AuthMiddleware())
	{
		fields.POST("", h.CreateField)
		fields.GET("/:id", h.GetField)
		fields.GET("/table/:table_id", h.GetFields)
		fields.PUT("/:id", h.UpdateField)
		fields.DELETE("/:id", h.DeleteField)
	}
}
