package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	PasswordHashTimesName   = "PASSWORD_HASH_TIMES"
	PasswordHashMemoryName  = "PASSWORD_HASH_MEMORY"
	PasswordHashThreadsName = "PASSWORD_HASH_THREADS"
	PasswordHashKeyLenName  = "PASSWORD_HASH_KEY_LENGTH"
	PasswordHashSaltLenName = "PASSWORD_HASH_SALT_LENGTH"
)

type PasswordHashParams struct {
	// Number of iterations
	Times uint32
	// Amount of memory to use, in KB
	Memory uint32
	// Parallelism, number of threads to use
	Threads uint8
	// Key length
	KeyLen uint32
	// Salt length
	SaltLen uint32
}

func DefaultPasswordHashParams() (*PasswordHashParams, error) {
	s := GetEnvWithDefault(PasswordHashTimesName, "1")
	times, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return nil, err
	}
	s = GetEnvWithDefault(PasswordHashMemoryName, "65536")
	memory, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return nil, err
	}
	s = GetEnvWithDefault(PasswordHashThreadsName, "4")
	threads, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return nil, err
	}
	s = GetEnvWithDefault(PasswordHashKeyLenName, "32")
	keyLen, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return nil, err
	}
	s = GetEnvWithDefault(PasswordHashSaltLenName, "16")
	saltLen, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return nil, err
	}
	return &PasswordHashParams{
		Times:   uint32(times),
		Memory:  uint32(memory),
		Threads: uint8(threads),
		KeyLen:  uint32(keyLen),
		SaltLen: uint32(saltLen),
	}, nil
}

// HashPassword generates a new password hash using the argon2id algorithm.
func HashPassword(password string, params PasswordHashParams) (string, error) {
	salt := make([]byte, params.SaltLen)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return "", err
	}
	hash := argon2.IDKey(
		[]byte(password), salt,
		params.Times, params.Memory, params.Threads, params.KeyLen,
	)

	// Encode the parameters, salt, and hash
	encoded := fmt.Sprintf(
		"$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		params.Memory, params.Times, params.Threads,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash),
	)

	return encoded, nil
}

// ComparePassword compares a password with an encoded hash to check if they
// match.
func ComparePassword(password, encodedHash string) (bool, error) {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return false, errors.New("invalid encoded hash format")
	}
	var memory uint32
	var iterations uint32
	var parallelism uint8
	n, err := fmt.Sscanf(
		parts[3], "m=%d,t=%d,p=%d", &memory, &iterations, &parallelism,
	)
	if err != nil {
		return false, err
	}
	if 3 != n {
		return false, errors.New("invalid encoded hash format")
	}
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}
	hash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}
	// Derive the key again using the same params
	derivedKey := argon2.IDKey(
		[]byte(password), salt, iterations, memory, parallelism,
		uint32(len(hash)),
	)
	// Compare using constant time
	if 1 == subtle.ConstantTimeCompare(hash, derivedKey) {
		return true, nil
	}
	return false, nil
}
