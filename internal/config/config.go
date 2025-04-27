package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using environment variables")
    }
}

func GetMongoURI() string {
    uri := os.Getenv("MONGO_URI")
    if uri == "" {
        log.Fatal("MONGO_URI not found in environment variables")
    }
    return uri
}
