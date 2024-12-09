package utils

import (
	"github.com/google/uuid"
)

// Generate a new UUID
func GenerateUUID() string {
	id := uuid.New()
	return id.String()
}
