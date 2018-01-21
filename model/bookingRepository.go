package model

import (
	"github.com/jinzhu/gorm"
)

func findBookForEventId(db* gorm.DB, eventId string) Booking{
	var booking Booking

	db.Where("event_id = ?", eventId).First(&booking)
	return booking
}

func bookEvent(db* gorm.DB, booking Booking){
	db.Create(&booking)
	db.Save(&booking)
}

func unBookEvent(db *gorm.DB, booking Booking){
	db.Delete(&booking)
}