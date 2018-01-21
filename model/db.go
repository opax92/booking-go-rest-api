package model

import (
	"github.com/jinzhu/gorm"
	"os"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitDb() *gorm.DB{
	db, err := gorm.Open("sqlite3", "./data.db")

	if err != nil {
		os.Exit(1)
	}


	db.AutoMigrate(&Event{}, &Booking{})

	return db
}
