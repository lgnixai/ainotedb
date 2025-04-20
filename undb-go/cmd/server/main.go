package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	spaceHandler "github.com/undb/undb-go/internal/space/handler"
	spaceModel "github.com/undb/undb-go/internal/space/model"
	spaceRepository "github.com/undb/undb-go/internal/space/repository"
	spaceRouter "github.com/undb/undb-go/internal/space/router"
	spaceService "github.com/undb/undb-go/internal/space/service"
	tableHandler "github.com/undb/undb-go/internal/table/handler"
	tableModel "github.com/undb/undb-go/internal/table/model"
	tableRepository "github.com/undb/undb-go/internal/table/repository"
	tableRouter "github.com/undb/undb-go/internal/table/router"
	tableService "github.com/undb/undb-go/internal/table/service"
	userHandler "github.com/undb/undb-go/internal/user/handler"
	userModel "github.com/undb/undb-go/internal/user/model"
	userRepository "github.com/undb/undb-go/internal/user/repository"
	userRouter "github.com/undb/undb-go/internal/user/router"
	userService "github.com/undb/undb-go/internal/user/service"
	viewModel "github.com/undb/undb-go/internal/view/model"
	viewRepository "github.com/undb/undb-go/internal/view/repository"
	viewRouter "github.com/undb/undb-go/internal/view/router"
	viewService "github.com/undb/undb-go/internal/view/service"

	fieldModel "github.com/undb/undb-go/internal/field/model"
	recordModel "github.com/undb/undb-go/internal/record/model"

	// Add realtime imports
	realtimeHandler "github.com/undb/undb-go/internal/realtime/handler"
	realtimeRouter "github.com/undb/undb-go/internal/realtime/router"
	realtimeService "github.com/undb/undb-go/internal/realtime/service"

	"github.com/undb/undb-go/pkg/config"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate models
	if err := db.AutoMigrate(
		&spaceModel.Space{},
		&spaceModel.Member{},
		&tableModel.Table{},

		&userModel.User{},
		&viewModel.View{},
		&recordModel.Record{},
		&fieldModel.Field{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize repositories
	spaceRepo := spaceRepository.NewSpaceRepository(db)
	tableRepo := tableRepository.NewTableRepository(db)
	userRepo := userRepository.NewUserRepository(db)
	viewRepo := viewRepository.NewViewRepository(db)

	// Initialize services
	spaceSvc := spaceService.NewSpaceService(spaceRepo)
	tableSvc := tableService.NewTableService(tableRepo)
	userSvc := userService.NewUserService(userRepo)
	viewSvc := viewService.NewViewService(viewRepo)
	// Initialize realtime service
	realtimeSvc := realtimeService.NewRealtimeService()
	go realtimeSvc.Run() // Start the realtime service in a goroutine

	// Initialize handlers
	spaceHdlr := spaceHandler.NewSpaceHandler(spaceSvc)
	tableHdlr := tableHandler.NewTableHandler(tableSvc)
	userHdlr := userHandler.NewUserHandler(userSvc)
	// Initialize realtime handler
	realtimeHdlr := realtimeHandler.NewRealtimeHandler() // TODO: Pass realtimeSvc if needed

	// Initialize router
	r := gin.Default()

	api := r.Group("/api")
	spaceRouter.RegisterSpaceRoutes(api, spaceHdlr)
	tableRouter.RegisterTableRoutes(api, tableHdlr)
	userRouter.RegisterUserRoutes(api, userHdlr)
	viewRouter.RegisterViewRoutes(api, viewSvc)
	realtimeRouter.RegisterRealtimeRoutes(api, realtimeHdlr)

	// Initialize view routes with view service
	viewSvc.Initialize()

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
