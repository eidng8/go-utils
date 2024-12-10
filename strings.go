package utils

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const (
	AlphaNum  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	Printable = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:,.<>?/`~"
)

func RandomAlphaNum(length int) (string, error) {
	return RandomString(AlphaNum, length)
}

func RandomPrintable(length int) (string, error) {
	return RandomString(Printable, length)
}

func RandomString(dict string, length int) (string, error) {
	result := make([]byte, length)
	charsetLength := big.NewInt(int64(len(dict)))
	for i := range result {
		num, err := randomInts(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		result[i] = dict[num.Int64()]
	}
	return string(result), nil
}

func StringIndexOfAny(s string, subs []string) int {
	for _, a := range subs {
		idx := strings.Index(s, a)
		if idx > -1 {
			return idx
		}
	}
	return -1
}

func StringContainsAny(s string, subs []string) bool {
	return StringIndexOfAny(s, subs) > -1
}
