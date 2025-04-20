
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
		
		tables := api.Group("/tables")
		{
			tables.GET("/:id", s.getTable)
			tables.GET("/space/:spaceId", s.listTables)
			tables.POST("", s.createTable)
			tables.PUT("/:id", s.updateTable)
			tables.DELETE("/:id", s.deleteTable)
			
			tables.POST("/:tableId/fields", s.addField)
			tables.PUT("/fields/:id", s.updateField)
			tables.DELETE("/fields/:id", s.deleteField)

			// Record endpoints
			tables.GET("/:tableId/records", s.listRecords)
			tables.GET("/:tableId/records/:id", s.getRecord)
			tables.POST("/:tableId/records", s.createRecord)
			tables.PUT("/:tableId/records/:id", s.updateRecord)
			tables.DELETE("/:tableId/records/:id", s.deleteRecord)
			
			// Bulk operations
			tables.POST("/:tableId/records/bulk", s.bulkCreateRecords)
			tables.PUT("/:tableId/records/bulk", s.bulkUpdateRecords)
			tables.DELETE("/:tableId/records/bulk", s.bulkDeleteRecords)

			// View endpoints
			tables.GET("/:tableId/views", s.listViews)
			tables.GET("/:tableId/views/:id", s.getView)
			tables.POST("/:tableId/views", s.createView)
			tables.PUT("/:tableId/views/:id", s.updateView)
			tables.DELETE("/:tableId/views/:id", s.deleteView)
			
			// View configuration
			tables.PUT("/:tableId/views/:id/fields", s.setViewFields)
			tables.PUT("/:tableId/views/:id/filter", s.setViewFilter)
			tables.PUT("/:tableId/views/:id/sort", s.setViewSort)
		}
	}
}

func (s *Server) getTable(c *gin.Context) {
	id := c.Param("id")
	table, err := s.table.GetTable(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, table)
}

func (s *Server) listTables(c *gin.Context) {
	spaceID := c.Param("spaceId")
	tables, err := s.table.ListTables(spaceID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, tables)
}

func (s *Server) createTable(c *gin.Context) {
	var table Table
	if err := c.ShouldBindJSON(&table); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	if err := s.table.CreateTable(&table); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, table)
}

func (s *Server) updateTable(c *gin.Context) {
	var table Table
	if err := c.ShouldBindJSON(&table); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	if err := s.table.UpdateTable(&table); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, table)
}

func (s *Server) deleteTable(c *gin.Context) {
	id := c.Param("id")
	if err := s.table.DeleteTable(id); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, nil)
}

func (s *Server) addField(c *gin.Context) {
	var field Field
	if err := c.ShouldBindJSON(&field); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	field.TableID = c.Param("tableId")
	if err := s.table.AddField(&field); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, field)
}

func (s *Server) updateField(c *gin.Context) {
	var field Field
	if err := c.ShouldBindJSON(&field); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	if err := s.table.UpdateField(&field); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, field)
}

func (s *Server) deleteField(c *gin.Context) {
	id := c.Param("id")
	if err := s.table.DeleteField(id); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, nil)
}

func (s *Server) Run() error {
	return s.router.Run(":5000")
}
