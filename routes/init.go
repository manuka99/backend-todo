package routes

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Initialize sets up the Gin router and starts the server
func Initialize() {
	// Create a new Gin router
	r := gin.Default()

	// Register all routes
	RegisterTaskRoutes(r)

	// Start the server
	log.Fatal(r.Run(":8080"))
}