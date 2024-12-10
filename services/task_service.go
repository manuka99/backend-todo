package services

import (
	"backend-todo/aws/dynamodb"
	"backend-todo/models"
	"backend-todo/utils"
	"log"
	"sync"
	"time"
)

const TaskCacheTTL = 5 * time.Minute

// Singleton variables for cache, fetch channel, and worker initialization.
var (
	taskCache *utils.Cache[string, models.Task]
	fetchChan chan utils.FetchRequest[string, models.Task]
	once      sync.Once
)

// InitializeTaskService initializes the cache, fetch channel, and starts the worker.
func InitializeTaskService() {
	once.Do(func() {
		// Initialize the cache with a TTL
		taskCache = utils.NewCache[string, models.Task]()
		fetchChan = make(chan utils.FetchRequest[string, models.Task])

		// Start the data-fetching worker for tasks (can be reused with other fetch functions).
		utils.StartDataFetcher(fetchChan, dynamodb.GetTask, taskCache, TaskCacheTTL)

		log.Println("Task service initialized with singleton access")
	})
}

// Accessor for the task cache (singleton access).
func GetTaskCache() *utils.Cache[string, models.Task] {
	return taskCache
}

// Accessor for the fetch channel (singleton access).
func GetFetchChannel() chan utils.FetchRequest[string, models.Task] {
	return fetchChan
}
