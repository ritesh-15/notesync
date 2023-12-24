package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDb() {
	var err error

	DB, err = gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Database connection established...")
}
