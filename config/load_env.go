package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT                      string
	HOST                      string
	DATABASE_URL              string
	VERIFICATION_TOKEN_SECRET string
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error while loading .env file", err)
		return
	}

	PORT = os.Getenv("PORT")
	HOST = os.Getenv("HOST")
	DATABASE_URL = os.Getenv("DATABASE_URL")
	VERIFICATION_TOKEN_SECRET = os.Getenv("VERIFICATION_TOKEN_SECRET")
}
