package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}

	log.Println("Loaded .env file successfully")
}
