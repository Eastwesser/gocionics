package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DbPassword string
	DbUser     string
	DbName     string
	DbPort     string
	DbHost     string
}

func NewConfig() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}

	return &Config{
		Port:       getEnv("PORT", "8080"),
		DbPassword: getEnv("DB_PASSWORD", "postgres"),
		DbUser:     getEnv("DB_USER", "postgres"),
		DbName:     getEnv("DB_NAME", "postgres"),
		DbPort:     getEnv("DB_PORT", "5432"),
		DbHost:     getEnv("DB_HOST", "db"),
	}

}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
