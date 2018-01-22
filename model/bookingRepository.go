package model

import (
	"github.com/jinzhu/gorm"
)

func findBookForEventId(db* gorm.DB, eventId uint64) Booking{
	var booking Booking

	db.Where("event_id = ?", eventId).First(&booking)
	return booking
}

func bookEvent(db* gorm.DB, booking Booking){
	db.Save(&booking)
}

func unBookEvent(db *gorm.DB, booking Booking){
	db.Delete(&booking)
}