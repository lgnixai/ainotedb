package router

import (
	"github.com/gin-gonic/gin"

	"github.com/undb/undb-go/internal/user/handler"
)

// RegisterUserRoutes registers all user-related routes
func RegisterRoutes(r *gin.RouterGroup, handler *handler.UserHandler) {
	users := r.Group("/users")
	{
		users.POST("/register", handler.Register)
		users.POST("/login", handler.Login)
		users.GET("/:id", handler.GetUser)
		users.PUT("/:id", handler.UpdateUser)
		users.DELETE("/:id", handler.DeleteUser)
	}
}
