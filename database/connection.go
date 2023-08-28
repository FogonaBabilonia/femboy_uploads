package database

import (
	"github.com/FogonaBabilonia/femboy_uploads/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	conn, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = conn
	DB.AutoMigrate(&models.User{})
}
