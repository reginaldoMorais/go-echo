package core

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func EnvMongoURI() string {
	getEnv()
	return os.Getenv("MONGO_URI")
}

func EnvMongoDatabase() string {
	getEnv()
	return os.Getenv("MONGO_DATABASE")
}
