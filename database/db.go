package database

import (
	"log"
	"os"
	"task5-pbi-btpns-holidmuhamadsalman/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Loading .env file")
	}

	dsn := os.Getenv("URL_DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{}, &models.Photo{}) 

	DB = db
}