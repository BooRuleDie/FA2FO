package env

import (
	"fmt"
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return val
}

func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	intVal, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return intVal
}

func MustGetString(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("Required environment variable %s is not set", key))
	}
	return val
}

func MustGetInt(key string) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("Required environment variable %s is not set", key))
	}

	intVal, err := strconv.Atoi(val)
	if err != nil {
		panic(fmt.Sprintf("Environment variable %s is not a valid integer: %v", key, err))
	}

	return intVal
}
