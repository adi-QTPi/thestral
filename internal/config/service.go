package config

import (
	"log"
	"os"
	"strconv"

	"github.com/adi-QTPi/thestral/internal/utils"
	"github.com/joho/godotenv"
)

const (
	DefaultAdminBind   = "0.0.0.0:80"
	DefaultProxyBind   = "0.0.0.0:7007"
	DefaultDatabaseURL = "host=azkaban user=user password=password dbname=thestral port=5433 sslmode=disable TimeZone=UTC"
	DefaultDebug       = false
)

type Env struct {
	PROXY_BIND   string `validate:"required,hostname_port"`
	ADMIN_BIND   string `validate:"required,hostname_port"`
	DATABASE_URL string `validate:"required"`
	DEBUG        bool   `validate:"required"`
}

// Loads the env variables and returns error in case of invalid variables
func LoadConfig() (*Env, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using Default Environment Variables")
	}

	env := &Env{
		PROXY_BIND:   getEnvString("PROXY_BIND", DefaultProxyBind),
		ADMIN_BIND:   getEnvString("ADMIN_BIND", DefaultAdminBind),
		DATABASE_URL: getEnvString("DATABASE_URL", DefaultDatabaseURL),
		DEBUG:        getEnvBool("DEBUG", DefaultDebug),
	}

	if err := utils.ValidateStruct(env); err != nil {
		return nil, err
	}

	return env, nil
}

func getEnvString(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvBool(key string, fallback bool) bool {
	valStr, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}

	value, err := strconv.ParseBool(valStr)
	if err != nil {
		return fallback
	}

	return value
}
