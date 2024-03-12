package utils

import (
	"fmt"
	"log"
	"os"

	"gds-onecv-asgt/models"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectToDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbURI := fmt.Sprintf("postgres://postgres:%s@database:5432/onecv-db", os.Getenv("DATABASE_PASSWORD"))

	var err error
	db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Student{}, &models.Teacher{})
}

func DB() *gorm.DB {
    return db
}
