package routes

import (
	"backend-todo/config"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// Initialize sets up the Gin router and starts the server
func Initialize() {
	// Create a new Gin router
	r := gin.Default()

	// Create a new group for the API
	api := r.Group("/api")

	// Register all routes
	RegisterTaskRoutes(api)

	// Health check route
	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	// Start the server
	log.Fatal(r.Run(fmt.Sprintf(":%s", config.PORT)))
}
