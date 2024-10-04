package main

import (
	"log"

	"github.com/androsyahreza/product-api/config"
	"github.com/androsyahreza/product-api/src/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		panic(err)
	}

	config.LoadEnv()

	config.InitializeDatabase()

	routes.ServerRoutes()
}
