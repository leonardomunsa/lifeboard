package storage

import (
	"fmt"
	"github.com/leonardomunsa/lifeboard/internal/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to DB: %v, err")
	}

	log.Println("Connected to database")

	err = DB.AutoMigrate(&models.Game{})
	if err != nil {
		log.Fatalf("Error auto-migrating database: %v", err)
	}
}
