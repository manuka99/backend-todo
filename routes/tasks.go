package routes

import (
	"backend-todo/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(r *gin.RouterGroup) {
	r.GET("/tasks", handlers.GetAllTasks)
	r.GET("/tasks/:id", handlers.GetTask)
	r.POST("/tasks", handlers.CreateTask)
}
