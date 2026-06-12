package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

func Load() Config {
	godotenv.Load()
	return Config{
		AppPort:    getEnvOrExit("APP_PORT", false),
		DBHost:     getEnvOrExit("DB_HOST", false),
		DBPort:     getEnvOrExit("DB_PORT", false),
		DBUser:     getEnvOrExit("DB_USER", false),
		DBPassword: getEnvOrExit("DB_PASSWORD", false),
		DBName:     getEnvOrExit("DB_NAME", false),
		DBSSLMode:  getEnvOrExit("DB_SSLMODE", false),
	}
}

func (c Config) DatabaseConnStr() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName, c.DBSSLMode)
}

func getEnvOrExit(key string, allowEmpty bool) string {
	if value, ok := os.LookupEnv(key); ok && (allowEmpty || len(value) > 0) {
		return value
	}
	panic(fmt.Sprintf("Env variable %s is not present", key))
}
