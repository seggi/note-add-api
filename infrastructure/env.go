package infrastructure

import (
	"log"

	"github.com/joho/godotenv"
)

// Loads environmesnt variables from .env file
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Unable to load .env file")
	}
}
