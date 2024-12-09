package utils

import (
	"backend-todo/models"
	"sync"
)

type TaskCache struct {
	tasks map[string]models.Task
	mu    sync.Mutex
}

func NewTaskCache() *TaskCache {
	return &TaskCache{tasks: make(map[string]models.Task)}
}

func (c *TaskCache) AddTask(task models.Task) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.tasks[task.Title] = task
}

func (c *TaskCache) GetAllTasks() []models.Task {
	c.mu.Lock()
	defer c.mu.Unlock()

	taskList := []models.Task{}
	for _, task := range c.tasks {
		taskList = append(taskList, task)
	}
	return taskList
}
