package handlers

import (
	"backend-todo/aws/dynamodb"
	"backend-todo/models"
	"backend-todo/services"
	"backend-todo/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetAllTasks: Handle GET /tasks
// Note that this handler fetches all tasks from DynamoDB and does not use the cache.
// Not caching all tasks is a design choice to avoid storing a large amount of data in memory.
func GetAllTasks(c *gin.Context) {
	tasks, err := dynamodb.FetchAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GetTask handles GET /tasks/:id to fetch a task from the cache or DynamoDB.
func GetTask(c *gin.Context) {
	id := c.Param("id")

	// Access the singleton cache and fetch channel
	taskCache := services.GetTaskCache()
	fetchChan := services.GetFetchChannel()

	// Check the cache first
	if data, found := taskCache.Get(id); found {
		c.JSON(http.StatusOK, data) // Return data from cache
		return
	}

	// If not found in cache, create a new FetchRequest
	req := utils.FetchRequest[string, models.Task]{
		ID:       id,
		Response: make(chan models.Task),
		Error:    make(chan error),
	}

	// Send the request to the worker for data fetching
	fetchChan <- req

	// Wait for a response or error using select
	select {
	case task := <-req.Response:
		// Successfully fetched the task, cache it and return the data
		taskCache.Set(id, task, services.TaskCacheTTL)
		c.JSON(http.StatusOK, task)

	case err := <-req.Error:
		// Handle error in case fetching fails
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch task: " + err.Error(),
		})

	case <-time.After(5 * time.Second):
		// Timeout for fetching data
		c.JSON(http.StatusGatewayTimeout, gin.H{
			"error": "Request timed out while fetching task",
		})
	}
}

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
	taskCache := services.GetTaskCache()
	taskCache.Set(task.Id, task, services.TaskCacheTTL)

	c.JSON(http.StatusCreated, task)
}
