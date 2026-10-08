package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port string
}

func LoadConfig() (*Config, error) {
	port := os.Getenv("PORT")

	if port == "" {
		return nil, fmt.Errorf("invalid environment variable/'s")
	}

	return &Config{
		Port: port,
	}, nil
}
