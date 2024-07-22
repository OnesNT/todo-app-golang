package config

import (
	"todo-app/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=ngokuang dbname=postgres sslmode=disable"
	database, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic("failed to connect to database")
	}

	database.AutoMigrate(&models.Task{})
	DB = database
}
