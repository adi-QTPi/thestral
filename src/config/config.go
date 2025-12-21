package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	PROXY_BIND     string
	ADMIN_BIND     string
	REDIS_HOST     string
	REDIS_PORT     string
	REDIS_PASSWORD string
}

func LoadConfig() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on System Environment Variables")
	}
	return &Env{
		PROXY_BIND:     getEnv("PROXY_BIND", "0.0.0.0:80"),
		ADMIN_BIND:     getEnv("ADMIN_BIND", "100.113.160.66:7007"),
		REDIS_HOST:     getEnv("REDIS_HOST", "localhost"),
		REDIS_PORT:     getEnv("REDIS_PORT", "6379"),
		REDIS_PASSWORD: getEnv("REDIS_PASSWORD", "secret-text"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
