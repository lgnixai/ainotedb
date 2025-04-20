package router

import (
	"github.com/gin-gonic/gin"
	"github.com/undb/undb-go/internal/field/handler"
)

// RegisterFieldRoutes 注册字段路由
func RegisterFieldRoutes(r *gin.RouterGroup, h *handler.FieldHandler) {
	fields := r.Group("/fields")
	{
		fields.POST("", h.CreateField)
		fields.GET("/:id", h.GetField)
		fields.GET("/table/:table_id", h.GetFields)
		fields.PUT("/:id", h.UpdateField)
		fields.DELETE("/:id", h.DeleteField)
	}
}
