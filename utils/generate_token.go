package utils

import (
	"crypto/rand"
	"math/big"
)

// generateRandomToken generates a random token of the specified length.
func GenerateRandomToken(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	token := make([]byte, length)
	for i := range token {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		token[i] = charset[num.Int64()]
	}
	return string(token), nil
}
