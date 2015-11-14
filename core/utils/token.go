package utils
import (
	"github.com/satori/go.uuid"
)

// GenerateToken generates a UUID v4 token
func GenerateToken(username string) string{
	return uuid.NewV4().String()
}

func GenerateUUID() string {
	return uuid.NewV4().String()
}
