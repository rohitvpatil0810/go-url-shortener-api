package utils

import (
	"math/rand"
	"strings"
	"time"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(chars)
	for i := 0; i < n; i++ {
		c := chars[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// RandomUsername generates a random username
func RandomUsername() string {
	return RandomString(7)
}

// RandomPassword generates a random password
func RandomPassword() string {
	return RandomString(10)
}
