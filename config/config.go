package config

import (
	"log"
	"os"

	"github.com/androsyz/product-api/src/models"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		panic(err)
	}
}

func InitializeDatabase() {
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")

	log.Printf("Connecting to database at %s:%s...", DbHost, DbPort)

	models.ConnectDatabase(DbUser, DbPassword, DbHost, DbPort, DbName)

	log.Printf("Database successfully connected!")
}
