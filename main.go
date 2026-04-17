package main

import (
	"log"

	"github.com/prakhar0009/go-todo/config"
	"github.com/prakhar0009/go-todo/database"
	"github.com/prakhar0009/go-todo/server"
)

func main() {
	config.LoadConfig()

	err := database.ConnectandMigrate(
		config.GetEnv("DB_HOST", "localhost"),
		config.GetEnv("DB_PORT", "5432"),
		config.GetEnv("DB_NAME", "todo_db"),
		config.GetEnv("DB_USER", "postgres"),
		config.GetEnv("DB_PASSWORD", "password"),
		database.SSLMode(config.GetEnv("DB_SSLMODE", "disable")),
	)
	if err != nil {
		log.Fatal(err)
	}

	r := server.NewServer()
	port := config.GetEnv("PORT", "8080")
	r.Run(":" + port)
}
