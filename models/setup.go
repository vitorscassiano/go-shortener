package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "shortener.db")

	if err != nil {
		panic("failed to connect database.")
	}

	database.AutoMigrate(&Shortener{})

	DB = database
}
