package dynamodb

import (
	"backend-todo/config"
	"backend-todo/models"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// FetchAllTasks retrieves all tasks from the DynamoDB table
func FetchAllTasks() ([]models.Task, error) {
	// Use the shared dynamoClient
	client := GetDynamoClient()

	// Create the Scan input
	input := &dynamodb.ScanInput{
		TableName: aws.String(config.TasksTable),
	}

	// Execute the Scan operation
	result, err := client.Scan(input)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tasks: %v", err)
	}

	// Map the results to Task models
	tasks := []models.Task{}
	for _, item := range result.Items {
		tasks = append(tasks, models.Task{
			Id:          *item["Id"].S,
			Title:       *item["Title"].S,
			Description: *item["Description"].S,
		})
	}
	return tasks, nil
}

// GetTask retrieves a task by ID from the DynamoDB table
func GetTask(id string) (models.Task, error) {
	// Use the shared dynamoClient
	client := GetDynamoClient()

	// Create the Query Input
	input := &dynamodb.QueryInput{
		TableName: aws.String(config.TasksTable),
		KeyConditions: map[string]*dynamodb.Condition{
			"Id": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(id),
					},
				},
			},
		},
	}

	// Execute the GetItem operation
	result, err := client.Query(input)
	if err != nil {
		return models.Task{}, fmt.Errorf("failed to get task: %v", err)
	}

	// Check if the task was found
	if result.Count == nil || *result.Count == 0 {
		return models.Task{}, fmt.Errorf("task not found")
	}

	// Map the result to a Task model
	task := models.Task{
		Id:          *result.Items[0]["Id"].S,
		Title:       *result.Items[0]["Title"].S,
		Description: *result.Items[0]["Description"].S,
	}

	return task, nil
}

// SaveTask saves a task to the DynamoDB table
func SaveTask(task models.Task) error {
	// Use the shared dynamoClient
	client := GetDynamoClient()

	// Create the PutItem input
	input := &dynamodb.PutItemInput{
		TableName: aws.String(config.TasksTable),
		Item: map[string]*dynamodb.AttributeValue{
			"Id":          {S: aws.String(task.Id)},
			"Title":       {S: aws.String(task.Title)},
			"Description": {S: aws.String(task.Description)},
		},
	}

	// Execute the PutItem operation
	_, err := client.PutItem(input)
	if err != nil {
		return fmt.Errorf("failed to save task: %v", err)
	}
	return nil
}
