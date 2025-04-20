
package main

import (
	"log"
	"retable/internal/app"
)

func main() {
	// Initialize database
	cfg := config.LoadConfig()
	db, err := database.NewDatabase(cfg.PostgresURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run auto migrations
	if err := db.AutoMigrate(); err != nil {
		log.Fatal("Failed to run database migrations:", err)
	}

	// Initialize and start server
	server := app.NewServer(db)
	if err := server.Run(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
