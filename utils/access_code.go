package utils

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"strconv"
)

func GenerateAccessCode() (string, error) {
	length, _ := strconv.Atoi(os.Getenv("ACCESS_CODE_LENGTH"))
	if length == 0 {
		length = 8
	}

	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}
