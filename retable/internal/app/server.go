
package app

import (
	"github.com/gin-gonic/gin"
	"retable/internal/config"
	"retable/internal/auth"
	"retable/internal/space"
	"retable/internal/cache"
	"retable/internal/realtime"
	"gorm.io/gorm"
)

type Server struct {
	config    *config.Config
	router    *gin.Engine
	db        *gorm.DB
	auth      *auth.AuthService
	space     *space.SpaceService
	cache     *cache.RedisCache
	websocket *realtime.WebSocketManager
}

func NewServer(db *gorm.DB) *Server {
	cfg := config.LoadConfig()
	router := gin.Default()
	
	authService := auth.NewAuthService(cfg.JWTSecret)
	spaceService := space.NewSpaceService(db)
	cacheService := cache.NewRedisCache(cfg.RedisURL)
	websocketManager := realtime.NewWebSocketManager()
	
	server := &Server{
		config:    cfg,
		router:    router,
		db:        db,
		auth:      authService,
		space:     spaceService,
		cache:     cacheService,
		websocket: websocketManager,
	}
	
	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	// Setup routes here
	api := s.router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			// Auth routes
		}
		
		spaces := api.Group("/spaces")
		{
			// Space routes
		}
		
		// Other routes
	}
}

func (s *Server) Run() error {
	return s.router.Run(":5000")
}
