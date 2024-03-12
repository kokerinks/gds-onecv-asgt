package utils

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"gds-onecv-asgt/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbURI := os.Getenv("DATABASE_URI")

	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Student{}, &models.Teacher{})
}
