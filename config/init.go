package config

import (
	"log"

	"github.com/joho/godotenv"
)

func Initialize() {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found. Proceeding with environment variables.")
	}
}
