package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	if err := godotenv.Load(".env"); err != nil {
        log.Fatalf("Error loading .env file")
    }
}