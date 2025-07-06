package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	GRPCHost string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, loading from environment variables")
	}

	log.Printf("Loaded API Gateway config")

	return &Config{
		Port:     getEnv("PORT", "8383"),
		GRPCHost: getEnv("GRPC_HOST", "localhost:50051"),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
