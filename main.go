package main

import (
	"log"

	"github.com/prakhar0009/go-todo/database"
	"github.com/prakhar0009/go-todo/server"
)

func main() {
	err := database.ConnectandMigrate(
		"localhost",
		"5432",
		"todo_db",
		"local",
		"local",
		database.SSLModeDisable,
	)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connected and migrated successfully")
	r := server.NewServer()

	log.Println("Server starting on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
