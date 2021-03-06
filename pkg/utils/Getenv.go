package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Getenv(key string) string {
	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}
	return os.Getenv(key)
}

func GetConnectURI() string {
	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}
	return "mongodb+srv://" + os.Getenv("mongoUsername") + ":" + os.Getenv("mongoPassword") + "@cluster0.4lioy.mongodb.net/eight-sup?retryWrites=true&w=majority"
}
