package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoDBURI string
}

// LoadConfig carrega as variáveis de ambiente e retorna uma configuração
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	mongoDBURI := os.Getenv("MONGO_URI")
	if mongoDBURI == "" {
		log.Fatal("MONGO_URI environment variable not set")
	}

	return &Config{
		MongoDBURI: mongoDBURI,
	}
}

// GetMongoURI retorna a URI do MongoDB
func GetMongoURI() string {
	return os.Getenv("MONGO_URI")
}
