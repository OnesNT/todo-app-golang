package config

import (
	"fmt"
	"os"
	"todo-app/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)
	DB, err = gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Printf("Error connecting to database: %s\n", err)
		os.Exit(-1)
	}
	fmt.Println("Successfully connected to the database")

	// AutoMigrate the Task model
	DB.AutoMigrate(&models.Task{})
}

func GetDB() *gorm.DB {
	return DB
}
