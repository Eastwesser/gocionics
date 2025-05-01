package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DB_password string
	DB_user     string
	DB_name     string
	DB_port     string
	DB_host     string
}

func NewConfig() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}

	return &Config{
		Port:        getEnv("PORT", "8080"),
		DB_password: getEnv("DB_PASSWORD", "postgres"),
		DB_user:     getEnv("DB_USER", "postgres"),
		DB_name:     getEnv("DB_NAME", "postgres"),
		DB_port:     getEnv("DB_PORT", "5432"),
		DB_host:     getEnv("DB_HOST", "db"),
	}

}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
