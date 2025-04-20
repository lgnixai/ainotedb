package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	spaceRouter "github.com/undb/undb-go/internal/space/router"

	"github.com/undb/undb-go/internal/config"
	fieldHandler "github.com/undb/undb-go/internal/field/handler"
	fieldModel "github.com/undb/undb-go/internal/field/model"
	fieldRouter "github.com/undb/undb-go/internal/field/router"
	fieldService "github.com/undb/undb-go/internal/field/service"
	db "github.com/undb/undb-go/internal/infrastructure/db"

	recordHandler "github.com/undb/undb-go/internal/record/handler"
	recordModel "github.com/undb/undb-go/internal/record/model"
	recordRouter "github.com/undb/undb-go/internal/record/router"
	recordService "github.com/undb/undb-go/internal/record/service"

	spaceHandler "github.com/undb/undb-go/internal/space/handler"
	spaceModel "github.com/undb/undb-go/internal/space/model"
	spaceService "github.com/undb/undb-go/internal/space/service"

	tableHandler "github.com/undb/undb-go/internal/table/handler"
	tableModel "github.com/undb/undb-go/internal/table/model"
	tableRouter "github.com/undb/undb-go/internal/table/router"
	tableService "github.com/undb/undb-go/internal/table/service"

	userHandler "github.com/undb/undb-go/internal/user/handler"
	userModel "github.com/undb/undb-go/internal/user/model"
	userRouter "github.com/undb/undb-go/internal/user/router"
	userService "github.com/undb/undb-go/internal/user/service"

	viewHandler "github.com/undb/undb-go/internal/view/handler"
	viewModel "github.com/undb/undb-go/internal/view/model"
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
	gormDB, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 初始化 Gin 路由
	r := gin.Default()
	// 设置CORS（允许前端跨域访问）
	//r.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"*"},
	//	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//}))

	//r.Use(func(c *gin.Context) {
	//	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	//	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	//	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	//	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
	//
	//	if c.Request.Method == "OPTIONS" {
	//		c.AbortWithStatus(204)
	//		return
	//	}
	//
	//	c.Next()
	//})

	r.Use(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin == "http://localhost:5173" || origin == "http://localhost:3000" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		}
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Auto migrate models
	if err := gormDB.AutoMigrate(
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
	spaceRepo := db.NewSpaceRepository(gormDB)
	spaceMemberRepo := db.NewMemberRepository(gormDB)
	tableRepo := db.NewTableRepository(gormDB)
	userRepo := db.NewUserRepository(gormDB)
	viewRepo := db.NewViewRepository(gormDB)
	recordRepo := db.NewRecordRepository(gormDB)
	fieldRepo := db.NewFieldRepository(gormDB)

	// Initialize services
	spaceSvc := spaceService.NewSpaceService(spaceRepo, spaceMemberRepo, gormDB)
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
	r = gin.Default()
	api := r.Group("/api")

	// Register public routes
	api.POST("/users/register", userHandler.Register)
	api.POST("/users/login", userHandler.Login)
	// 可选：公开空间列表（如有）
	// api.GET("/spaces/public", spaceHandler.ListPublicSpaces)

	// Register模块路由，权限控制交由各自RegisterRoutes实现
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
