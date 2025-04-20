
package main

import (
	"log"
	"retable/internal/app"
)

func main() {
	server := app.NewServer()
	if err := server.Run(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
