package config

import (
	"log"
	"os"
)

type Config struct {
	ChatGPTAPIKey string
}

func LoadConfig() *Config {
	apiKey := os.Getenv("CHATGPT_API_KEY")
	if apiKey == "" {
		log.Fatal("CHATGPT_API_KEY environment variable is required")
	}

	return &Config{
		ChatGPTAPIKey: apiKey,
	}
}
