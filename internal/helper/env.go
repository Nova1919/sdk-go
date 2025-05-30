package helper

import (
	"os"
	"strconv"
)

// GetString returns the env variable for the given key
// and falls back to the given defaultValue if not set
func GetString(key, defaultValue string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}
	return defaultValue
}

// GetInt returns the env variable (parsed as integer) for
// the given key and falls back to the given defaultValue if not set
func GetInt(key string, defaultValue int) int {
	v, ok := os.LookupEnv(key)
	if ok {
		value, err := strconv.Atoi(v)
		if err != nil {
			return defaultValue
		}
		return value
	}
	return defaultValue
}

// GetInt64 returns the env variable (parsed as integer) for
// the given key and falls back to the given defaultValue if not set
func GetInt64(key string, defaultValue int64) int64 {
	v, ok := os.LookupEnv(key)
	if ok {
		value, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return defaultValue
		}
		return value
	}
	return defaultValue
}

// GetFloat64 returns the env variable (parsed as float64) for
// the given key and falls back to the given defaultValue if not set
func GetFloat64(key string, defaultValue float64) float64 {
	v, ok := os.LookupEnv(key)
	if ok {
		value, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return defaultValue
		}
		return value
	}
	return defaultValue
}

// GetBool returns the env variable (parsed as bool) for
// the given key and falls back to the given defaultValue if not set
func GetBool(key string, defaultValue bool) bool {
	v, ok := os.LookupEnv(key)
	if ok {
		value, err := strconv.ParseBool(v)
		if err != nil {
			return defaultValue
		}
		return value
	}
	return defaultValue
}
