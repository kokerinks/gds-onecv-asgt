package utils

import (
	"fmt"
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

	dbURI := fmt.Sprintf("postgres://postgres:%s@database:5432/onecv-db", os.Getenv("DATABASE_PASSWORD"))

	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Student{}, &models.Teacher{})
}
