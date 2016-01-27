package utils

import "math/rand"

const options = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RanStringBytesRmdr is from http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = options[rand.Int63() % int64(len(options))]
	}
	return string(b)
}