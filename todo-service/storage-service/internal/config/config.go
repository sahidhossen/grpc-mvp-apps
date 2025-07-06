package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// All configuration for storage service
type Config struct {
	GRPCPort string
	DBPath   string
}

// LoadConfig loads the configurations
func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, loading from environemtn variables")
	}

	log.Printf("Loaded storage service config")

	return &Config{
		GRPCPort: getEnv("GRPC_PORT", "50051"),
		DBPath:   getEnv("DB_PATH", "./data/todo.db"), // Default path
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
