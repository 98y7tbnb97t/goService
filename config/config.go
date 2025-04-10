package config

import (
	"os"
	"strconv"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type JWTConfig struct {
	Secret          string
	RefreshSecret   string
	AccessDuration  int
	RefreshDuration int
}

type Config struct {
	DB  DBConfig
	JWT JWTConfig
}

// parseInt safely converts a string to an integer with a default value of 0
func parseInt(value string) int {
	if value == "" {
		return 0
	}

	result, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}

	return result
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	return &Config{
		DB: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		JWT: JWTConfig{
			Secret:          os.Getenv("JWT_SECRET"),
			RefreshSecret:   os.Getenv("JWT_REFRESH_SECRET"),
			AccessDuration:  parseInt(os.Getenv("JWT_ACCESS_DURATION")),
			RefreshDuration: parseInt(os.Getenv("JWT_REFRESH_DURATION")),
		},
	}, nil
}
