package config

import (
	"log"

	"github.com/joho/godotenv"
)

func Initialize() {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
