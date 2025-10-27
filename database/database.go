package database

import (
	"log"

	"github.com/emmanuelYohore/api-rest-crud-go/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	log.Println("Connection database success")
	DB.AutoMigrate(&models.Product{})

}
