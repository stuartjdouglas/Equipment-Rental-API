package utils

import (
	"crypto/sha512"
	"encoding/hex"
)

// Sha512Me Returns an SHA512 encrypted string
func Sha512Me(data []byte) string {
	bytes := sha512.Sum512(data)
	return hex.EncodeToString(bytes[:])
}
