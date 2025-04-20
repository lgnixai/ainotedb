
package app

import (
	"github.com/gin-gonic/gin"
	"retable/internal/config"
	"retable/internal/core/domain"
)

type Server struct {
	config *config.Config
	router *gin.Engine
}

func NewServer() *Server {
	cfg := config.LoadConfig()
	router := gin.Default()
	
	return &Server{
		config: cfg,
		router: router,
	}
}

func (s *Server) Run() error {
	return s.router.Run(":5000")
}
