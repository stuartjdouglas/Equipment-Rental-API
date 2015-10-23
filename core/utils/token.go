package utils
import (
	"github.com/satori/go.uuid"
)

// Generates a UUID v4 token

func GenerateToken(username string) string{
	return uuid.NewV4().String()
}
