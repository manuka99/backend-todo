package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
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
}
