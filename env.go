package utils

import "os"

// GetEnvWithDefault returns the value of the environment variable named by
// the key. It is guaranteed to return the default value if the environment
// variable is not found or is empty.
func GetEnvWithDefault(key, defaultValue string) string {
	return GetEnvWithDefaultNP(key, defaultValue, false)
}

// GetEnvWithDefaultNP returns the value of the environment variable named by
// the key. It is guaranteed to return the default value if the environment
// variable is not found. If the environment variable is found, it is guaranteed
// to return the default value if the environment variable is empty.
func GetEnvWithDefaultNP(key, defaultValue string, nonEmpty bool) string {
	val, found := os.LookupEnv(key)
	if !found {
		return defaultValue
	}
	if nonEmpty && "" == val {
		return defaultValue
	}
	return val
}
