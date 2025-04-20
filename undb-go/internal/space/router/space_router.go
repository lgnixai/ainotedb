package router

import (
	"github.com/gin-gonic/gin"

	"github.com/undb/undb-go/internal/space/handler"
	`github.com/undb/undb-go/internal/user/middleware`
)

// RegisterSpaceRoutes 注册空间路由
func RegisterRoutes(r *gin.RouterGroup, h *handler.SpaceHandler) {
	spaces := r.Group("/spaces")
	spaces.Use(middleware.AuthMiddleware())

	{
		spaces.POST("", h.CreateSpace)
		spaces.GET("", h.ListSpaces)
		spaces.GET("/:id", h.GetSpace)
		spaces.PUT("/:id", h.UpdateSpace)
		spaces.DELETE("/:id", h.DeleteSpace)

		spaces.POST("/:id/members", h.AddMember)
		spaces.DELETE("/:id/members/:user_id", h.RemoveMember)
		spaces.PUT("/:id/members/:user_id/role", h.UpdateMemberRole)
		spaces.GET("/:id/members", h.GetSpaceMembers)
	}
}
