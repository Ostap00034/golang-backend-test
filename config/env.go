package config

import (
	"log"
	"os"

	"github.com/lpernett/godotenv"
)

var ConnStr = initConfig()

func initConfig() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return getEnv("CONNECTION_STRING", "user=postgres password=roottoor dbname=ecom")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
