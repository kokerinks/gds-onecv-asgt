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

func ConnectToDB(isTesting bool) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	var dbURI string
	if !isTesting {
		dbURI = fmt.Sprintf("postgres://postgres:%s@database:5432/onecv-db", os.Getenv("DATABASE_PASSWORD"))
	} else {
		dbURI = "postgres://postgres@localhost:5432/onecv-db"
	}

	var err error
	db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if sqlDB, err := db.DB(); err != nil {
		log.Fatalf("Error getting underlying database connection: %v", err)
	} else if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	db.AutoMigrate(&models.Student{}, &models.Teacher{})

	log.Printf("Database connection established.")
}

func DB() *gorm.DB {
	return db
}
