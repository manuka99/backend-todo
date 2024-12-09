package aws

import (
	"sync"

	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	awsSession *session.Session
	once       sync.Once
)

// GetAWSSession initializes (once) and returns a shared AWS session
func GetAWSSession() *session.Session {
	once.Do(func() {
		// Create a new session with shared configuration
		awsSession = session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
	})
	return awsSession
}
