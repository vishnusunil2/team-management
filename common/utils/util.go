package utils

import (
	"crypto/rand"
	"golang.org/x/crypto/bcrypt"
	"math/big"
	"strings"
)

var AlphaNumeric = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func GenerateRandomString(chars []rune, length int, prefix ...string) (string, error) {
	var result strings.Builder
	if len(prefix) > 0 && prefix[0] != "" {
		result.WriteString(prefix[0])
	}
	charCount := big.NewInt(int64(len(chars)))
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, charCount)
		if err != nil {
			return "", err
		}
		result.WriteRune(chars[index.Int64()])
	}
	return result.String(), nil
}
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
