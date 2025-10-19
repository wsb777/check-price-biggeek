package config

import (
	"log"
	"os"
)

type Config struct {
	BotToken string
}

func NewConfig() *Config {
	return &Config{
		BotToken: getEnv("BOT_TOKEN"),
	}
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("[ERROR] Empty env: %s", key)
	}
	return value
}
