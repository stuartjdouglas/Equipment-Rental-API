package utils
import (
	"crypto/sha512"
	"encoding/hex"
)

func ShaSum(data []byte) string {
	bytes := sha512.Sum512(data)
	return hex.EncodeToString(bytes[:])
}


func EncyptString(data string, filetype string) string {
	butes := sha512.Sum512([]byte(data))
	return hex.EncodeToString(butes[:]) + filetype
}