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
	SENDGRID_API_KEY          string
	CLIENT_URL                string
	SENDGRID_EMAIL            string
	ACCESS_TOKEN_SECRET       string
	REFRESH_TOKEN_SECRET      string
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
	SENDGRID_API_KEY = os.Getenv("SENDGRID_API_KEY")
	CLIENT_URL = os.Getenv("CLIENT_URL")
	SENDGRID_EMAIL = os.Getenv("SENDGRID_EMAIL")
	ACCESS_TOKEN_SECRET = os.Getenv("ACCESS_TOKEN_SECRET")
	REFRESH_TOKEN_SECRET = os.Getenv("REFRESH_TOKEN_SECRET")
}
