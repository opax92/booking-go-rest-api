package model

import (
	"github.com/jinzhu/gorm"
)

func findAllEvents(db* gorm.DB) []Event{
	var events[] Event

	db.Find(&events)
	return events
}

func deleteEvent(db* gorm.DB, event* Event){
	db.Delete(event)
}

func findEventById(db* gorm.DB, id uint64) *Event{
	var event Event

	db.First(&event, id)
	return &event
}

func findEventByName(db* gorm.DB, name string) Event{
	var event Event

	db.Where("event_name = ?", name).First(&event)
	return event
}

func createEvent(db* gorm.DB, event* Event){
	db.Create(event)
	db.Save(event)
}