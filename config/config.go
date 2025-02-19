package config

import (
	"os"

	"github.com/DanikDaraboz/StoreProject/pkg/logger"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	MongoURI   string
}

// LoadConfig reads from .env file
func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		logger.WarnLogger.Println("No .env file found, using system environment variables")
	}

	return Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		MongoURI:   getEnv("MONGO_URI", "mongodb+srv://danchik:denchik2006@cluster0.ed6vt.mongodb.net/"),
	}
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		logger.InfoLogger.Println("Using default value for", key)
		return fallback
	}
	return value
}
