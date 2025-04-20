package router

import (
	"github.com/gin-gonic/gin"
	"github.com/undb/undb-go/internal/view/handler"
	"github.com/undb/undb-go/internal/view/service"
)

// RegisterViewRoutes 注册视图相关路由
func RegisterViewRoutes(r *gin.RouterGroup, svc service.ViewService) {
	h := handler.NewViewHandler(svc)

	views := r.Group("/views")
	{
		views.POST("", h.CreateView)
		views.GET(":id", h.GetView)
		views.GET("/table/:tableId", h.GetViews)
		views.PUT(":id", h.UpdateView)
		views.DELETE(":id", h.DeleteView)
		views.PUT(":id/config", h.UpdateViewConfig)
	}
}
