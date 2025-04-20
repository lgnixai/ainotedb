package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/undb/undb-go/internal/config"
	fieldHandler "github.com/undb/undb-go/internal/field/handler"
	fieldModel "github.com/undb/undb-go/internal/field/model"
	fieldRepository "github.com/undb/undb-go/internal/field/repository"
	fieldRouter "github.com/undb/undb-go/internal/field/router"
	fieldService "github.com/undb/undb-go/internal/field/service"

	recordHandler "github.com/undb/undb-go/internal/record/handler"

	recordModel "github.com/undb/undb-go/internal/record/model"
	recordRepository "github.com/undb/undb-go/internal/record/repository"
	recordRouter "github.com/undb/undb-go/internal/record/router"
	recordService "github.com/undb/undb-go/internal/record/service"

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

	viewHandler "github.com/undb/undb-go/internal/view/handler"
	viewModel "github.com/undb/undb-go/internal/view/model"
	viewRepository "github.com/undb/undb-go/internal/view/repository"
	viewRouter "github.com/undb/undb-go/internal/view/router"
	viewService "github.com/undb/undb-go/internal/view/service"

	realtimeHandler "github.com/undb/undb-go/internal/realtime/handler"
	realtimeRouter "github.com/undb/undb-go/internal/realtime/router"
	realtimeService "github.com/undb/undb-go/internal/realtime/service"
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
		&spaceModel.SpaceMember{},
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
	spaceMemberRepo := spaceRepository.NewMemberRepository(db)
	tableRepo := tableRepository.NewTableRepository(db)
	userRepo := userRepository.NewUserRepository(db)
	viewRepo := viewRepository.NewViewRepository(db)
	recordRepo := recordRepository.NewRecordRepository(db)
	fieldRepo := fieldRepository.NewFieldRepository(db)

	// Initialize services
	spaceSvc := spaceService.NewSpaceService(spaceRepo, spaceMemberRepo, db)
	tableSvc := tableService.NewTableService(tableRepo)
	userSvc := userService.NewUserService(userRepo)
	viewSvc := viewService.NewViewService(viewRepo)
	recordSvc := recordService.NewRecordService(recordRepo)
	fieldSvc := fieldService.NewFieldService(fieldRepo)
	realtimeSvc := realtimeService.NewRealtimeService()

	// Initialize handlers
	spaceHandler := spaceHandler.NewSpaceHandler(spaceSvc)
	tableHandler := tableHandler.NewTableHandler(tableSvc)
	userHandler := userHandler.NewUserHandler(userSvc)
	viewHandler := viewHandler.NewViewHandler(viewSvc)
	recordHandler := recordHandler.NewRecordHandler(recordSvc)
	fieldHandler := fieldHandler.NewFieldHandler(fieldSvc)
	realtimeHandler := realtimeHandler.NewRealtimeHandler(realtimeSvc)

	// Initialize Gin router
	r := gin.Default()
	api := r.Group("/api")

	// Register routes
	spaceRouter.RegisterRoutes(api, spaceHandler)
	tableRouter.RegisterRoutes(api, tableHandler)
	userRouter.RegisterRoutes(api, userHandler)
	viewRouter.RegisterRoutes(api, viewHandler)
	recordRouter.RegisterRoutes(api, recordHandler)
	fieldRouter.RegisterRoutes(api, fieldHandler)
	realtimeRouter.RegisterRoutes(r, realtimeHandler)

	// Start server
	addr := fmt.Sprintf("0.0.0.0:%s", cfg.ServerPort)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
