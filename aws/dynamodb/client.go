package dynamodb

import (
	"backend-todo/aws"
	"sync"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	dynamoClient     *dynamodb.DynamoDB
	onceDynamoClient sync.Once
)

// GetDynamoClient initializes (once) and returns a shared DynamoDB client
func GetDynamoClient() *dynamodb.DynamoDB {
	onceDynamoClient.Do(func() {
		// Get the shared AWS session
		sess := aws.GetAWSSession()
		// Initialize DynamoDB client using the shared session
		dynamoClient = dynamodb.New(sess)
	})

	return dynamoClient
}
