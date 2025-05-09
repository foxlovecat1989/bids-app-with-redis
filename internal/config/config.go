package config

import (
	"os"
	"strconv"
)

// Redis holds Redis-specific configuration
type Redis struct {
	Host     string
	Port     int
	DB       int
	Password string
}

// Config holds all the configuration for the application
type Config struct {
	ServerPort int
	Redis      *Redis
}

// NewConfig creates a new Config with values from environment variables
// or defaults if environment variables are not set
func NewConfig() *Config {
	redis := &Redis{
		Host:     getEnv("REDIS_HOST", "localhost"),
		Port:     getEnvAsInt("REDIS_PORT", 6379),
		DB:       getEnvAsInt("REDIS_DB", 0),
		Password: getEnv("REDIS_PASSWORD", ""),
	}

	return &Config{
		ServerPort: getEnvAsInt("SERVER_PORT", 8080),
		Redis:      redis,
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
