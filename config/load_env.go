package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT string
	HOST string
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error while loading .env file", err)
		return
	}

	PORT = os.Getenv("PORT")
	HOST = os.Getenv("HOST")
}
