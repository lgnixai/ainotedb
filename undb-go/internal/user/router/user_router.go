package router

import (
	"github.com/gin-gonic/gin"

	"github.com/undb/undb-go/internal/user/handler"
)

// RegisterUserRoutes registers all user-related routes
import "github.com/undb/undb-go/internal/user/middleware"

func RegisterRoutes(r *gin.RouterGroup, handler *handler.UserHandler) {
	// 只注册需要登录的用户操作，公开路由由 main.go 注册，避免重复注册导致 panic
	authUsers := r.Group("/users")
	authUsers.Use(middleware.AuthMiddleware())
	{
		authUsers.GET("/:id", handler.GetUser)
		authUsers.PUT("/:id", handler.UpdateUser)
		authUsers.DELETE("/:id", handler.DeleteUser)
	}
}
