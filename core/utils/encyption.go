package utils
import (
	"crypto/sha512"
	"encoding/hex"
)

func ShaSum(data []byte) string {
	bytes := sha512.Sum512(data)
	return hex.EncodeToString(bytes[:])
}