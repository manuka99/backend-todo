package handlers

import (
	"backend-todo/aws/dynamodb"
	"backend-todo/models"
	"backend-todo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// In-memory cache
var taskCache = utils.NewTaskCache()

// CreateTask: Handle POST /tasks
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.Id = utils.GenerateUUID()

	// Save to DynamoDB
	if err := dynamodb.SaveTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update cache
	taskCache.AddTask(task)

	c.JSON(http.StatusCreated, task)
}

// GetAllTasks: Handle GET /tasks
func GetAllTasks(c *gin.Context) {
	tasks := taskCache.GetAllTasks()
	c.JSON(http.StatusOK, tasks)
}
