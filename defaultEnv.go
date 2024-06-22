package goli

import (
	"fmt"
	"os"
	"strconv"
)

// DefaultEnv returns either the env var at key or the default value
// Note: An empty-string env var is counted as missing. Use DefaultEnvEmpty to include empty env vars
func DefaultEnv(key string, defaultValue string) string {
	v := DefaultEnvEmpty(key, defaultValue)
	if v == "" {
		v = defaultValue
	}
	return v
}

// DefaultEnvEmpty returns an env var or default value. Empty env vars are considered found.
func DefaultEnvEmpty(key string, defaultValue string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return defaultValue
}

// DefaultEnvInt returns an env var converted to an int.
func DefaultEnvInt(key string, defaultValue int) (int, error) {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue, nil
	}
	intV, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("converting env:%s: %w", key, err)
	}
	return intV, nil
}
