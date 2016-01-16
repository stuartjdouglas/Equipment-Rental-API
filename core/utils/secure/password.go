package secure

import (
	"crypto/sha512"
	"io"
	"crypto/rand"
	"log"
)

const saltSize = 64

func SaltPassword(password string) []byte {
	buffer := make([]byte, saltSize, saltSize + sha512.Size)
	_, err := io.ReadFull(rand.Reader, buffer)

	if err != nil {
		log.Println("failed to random read: %s", err)
	}

	hash := sha512.New()
	hash.Write(buffer)
	hash.Write([]byte(password))
	return hash.Sum(buffer)
}
