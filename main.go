package main

import (
	"log"
	"os"

	"github.com/androsyahreza/product-api/src/models"
	"github.com/androsyahreza/product-api/src/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		panic(err)
	}

	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")

	models.ConnectDatabase(DbUser, DbPassword, DbHost, DbPort, DbName)

	routes.ServerRoutes()
}
