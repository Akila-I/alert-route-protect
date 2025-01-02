// File: config/config.go
package config

import (
	"os"
)

type Config struct {
	Port         string
	ClientID     string
	ClientSecret string
}

func LoadConfig() *Config {
	return &Config{
		Port:         getEnvOrDefault("PORT", "8080"),
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
