package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	// Port is the server port
	PORT = "8080"
)

func Initialize() {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found. Proceeding with environment variables.")
	}

	// Set environment variables
	PORT = os.Getenv("PORT")
}
