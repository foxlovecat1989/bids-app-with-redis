package config

import (
	"os"
	"strconv"
)

// Config holds all the configuration for the application
type Config struct {
	ServerPort  int
	RedisHost   string
	RedisPort   int
	RedisDB     int
	RedisPasswd string
}

// NewConfig creates a new Config with values from environment variables
// or defaults if environment variables are not set
func NewConfig() *Config {
	return &Config{
		ServerPort:  getEnvAsInt("SERVER_PORT", 8080),
		RedisHost:   getEnv("REDIS_HOST", "localhost"),
		RedisPort:   getEnvAsInt("REDIS_PORT", 6379),
		RedisDB:     getEnvAsInt("REDIS_DB", 0),
		RedisPasswd: getEnv("REDIS_PASSWORD", ""),
	}
}

// Helper functions to get environment variables
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
