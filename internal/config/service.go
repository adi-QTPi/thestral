package config

import (
	"log"
	"os"
	"strconv"

	"github.com/adi-QTPi/thestral/internal/utils"
	"github.com/joho/godotenv"
)

// redundant, as validate tags ensure user specified vals
const (
	DefaultAdminBind    = "0.0.0.0:7007"
	DefaultProxyBind    = "0.0.0.0:80"
	DefaultProxySSLBind = "0.0.0.0:443"
	DefaultDatabaseURL  = "host=azkaban user=user password=password dbname=thestral port=5433 sslmode=disable TimeZone=UTC"
	DefaultDebug        = false
	DefaultEnableTLS    = false
	DefaultReqPerSec    = 1
	DefaultBurst        = 5
)

type Env struct {
	PROXY_BIND             string `validate:"required,hostname_port"`
	PROXY_SSL_BIND         string `validate:"required,hostname_port"`
	ADMIN_BIND             string `validate:"required,hostname_port"`
	DATABASE_URL           string `validate:"required"`
	DEBUG                  bool
	ENABLE_TLS             bool
	RATE_LIMIT_REQ_PER_SEC int
	RATE_LIMIT_BURST       int
}

// Loads the env variables and returns error in case of invalid variables
func LoadConfig() (*Env, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using Default Environment Variables")
	}

	env := &Env{
		PROXY_BIND:             getEnvString("PROXY_BIND", DefaultProxyBind),
		PROXY_SSL_BIND:         getEnvString("PROXY_SSL_BIND", DefaultProxySSLBind),
		ADMIN_BIND:             getEnvString("ADMIN_BIND", DefaultAdminBind),
		DATABASE_URL:           getEnvString("DATABASE_URL", DefaultDatabaseURL),
		DEBUG:                  getEnvBool("DEBUG", DefaultDebug),
		ENABLE_TLS:             getEnvBool("ENABLE_TLS", DefaultEnableTLS),
		RATE_LIMIT_REQ_PER_SEC: getEnvInt("RATE_LIMIT_REQ_PER_SEC", DefaultReqPerSec),
		RATE_LIMIT_BURST:       getEnvInt("RATE_LIMIT_BURST", DefaultBurst),
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

func getEnvInt(key string, fallback int) int {
	valStr, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}

	value, err := strconv.Atoi(valStr)
	if err != nil {
		return fallback
	}

	return value
}
