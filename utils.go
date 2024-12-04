package utils

import (
	"os"
)

// GetEnvWithDefault returns the value of the environment variable named by
// the key. It is guaranteed to return the default value if the environment
// variable is not found or is empty.
func GetEnvWithDefault(key, defaultValue string) string {
	return GetEnv(key, defaultValue, false)
}

// GetEnvWithDefaultNE returns the value of the environment variable named by
// the key. It is guaranteed to return the default value if the environment
// variable is not found. If the environment variable is found, it is guaranteed
// to return the default value if the environment variable is empty.
func GetEnvWithDefaultNE(key, defaultValue string) string {
	return GetEnv(key, defaultValue, true)
}

// GetEnv returns the value of the environment variable named by
// the key. It is guaranteed to return the `defaultValue` if the environment
// variable is not found. If the environment variable is found, it is guaranteed
// to return the `defaultValue` if the environment variable is empty and
// `nonEmpty` is `true`.
func GetEnv(key, defaultValue string, nonEmpty bool) string {
	val, found := os.LookupEnv(key)
	if !found {
		return defaultValue
	}
	if nonEmpty && "" == val {
		return defaultValue
	}
	return val
}

// PanicIfError panics if the error is not nil.
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
