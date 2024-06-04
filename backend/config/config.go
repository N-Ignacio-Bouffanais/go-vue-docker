package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func GetToken() string {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatalf("TOKEN environment variable not set")
	}
	return token
}
