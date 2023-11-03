package uuidx

import "github.com/google/uuid"

// UUID generates a non-hermetic uuid
func UUID() string {
	return uuid.NewString()
}
