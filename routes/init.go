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

	RegisterRoutes(r)

	// Start the server
	log.Fatal(r.Run(fmt.Sprintf(":%s", config.PORT)))
}
