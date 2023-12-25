package main

import (
	"log"

	"github.com/ritesh-15/notesync-backend/config"
	"github.com/ritesh-15/notesync-backend/models"
)

func init() {
	config.LoadEnv()
	config.InitDb()
}

func main() {
	log.Println("running migrations...")

	err := config.DB.AutoMigrate(&models.User{}, &models.Session{}, &models.Document{}, &models.Folder{}, &models.Workspace{}, &models.Collaborator{})

	if err != nil {
		log.Fatal("migration failed", err)
	}

	log.Println("migration succesfully done")
}