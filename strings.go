package utils

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"
	"strings"
)

const (
	Printable = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:,.<>?/`~"
)

func RandomAlphaNum(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}

func RandomPrintable(length int) (string, error) {
	result := make([]byte, length)
	charsetLength := big.NewInt(int64(len(Printable)))
	for i := range result {
		num, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		result[i] = Printable[num.Int64()]
	}
	return string(result), nil
}

func StringIndexOfAny(s string, subs []string) int {
	for i, a := range subs {
		if strings.Contains(s, a) {
			return i
		}
	}
	return -1
}

func StringContainsAny(s string, subs []string) bool {
	return StringIndexOfAny(s, subs) > -1
}
