package memory

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateRandomString(length int) string {
	s := make([]byte, length)
	for i := range s {
		s[i] = letters[rand.Int63n(int64(length))]
	}

	return string(s)
}
