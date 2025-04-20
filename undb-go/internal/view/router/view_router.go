package router

import (
	"github.com/gin-gonic/gin"

	"github.com/undb/undb-go/internal/view/handler"
)

// RegisterViewRoutes 注册视图相关路由
import "github.com/undb/undb-go/internal/user/middleware"

func RegisterRoutes(r *gin.RouterGroup, h *handler.ViewHandler) {
	views := r.Group("/views")
	views.Use(middleware.AuthMiddleware())
	{
		views.POST("", h.CreateView)
		views.GET(":id", h.GetView)
		views.GET("/table/:tableId", h.GetViews)
		views.PUT(":id", h.UpdateView)
		views.DELETE(":id", h.DeleteView)
		views.PUT(":id/config", h.UpdateViewConfig)
	}
}
