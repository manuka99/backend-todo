package routes

import (
	"backend-todo/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(r *gin.Engine) {
	r.POST("/tasks", handlers.CreateTask)
	r.GET("/tasks", handlers.GetAllTasks)
}
