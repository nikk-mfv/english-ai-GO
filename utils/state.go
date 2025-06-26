package utils

import (
	"math/rand"
)

func GenerateState() string {
	// Generate a random state string for OAuth2
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	state := make([]byte, 32)
	for i := range state {
		state[i] = charset[rand.Intn(len(charset))]
	}
	return string(state)
}
