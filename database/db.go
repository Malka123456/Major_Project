package database

import (
	"fmt"
	"learning-backend/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=Aa09!@#$ dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Kolkata"

	var err error
	DB, err = gorm.Open( postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to DB", err)
	}

	fmt.Println("Database connected successfully")

	DB.AutoMigrate(&models.User{})


}