package model

import "github.com/jinzhu/gorm"

type EventView struct{
	Id        uint64
	EventName string
	BookedBy string
}

func getAllEventsView(db* gorm.DB) []EventView{
	var allEvents = findAllEvents(db)
	var eventView []EventView

	for i := 0; i < len(allEvents); i++{
		isEventBooked, bookedEvent := eventIsBooked(db, allEvents[i])

		if isEventBooked{
			eventView = append(eventView, eventViewWithBookedByValue(allEvents[i], bookedEvent.BookedBy))
		}else{
			eventView = append(eventView, eventViewWithBookedByValue(allEvents[i], ""))
		}
	}

	return eventView
}

func eventViewWithBookedByValue(event Event, bookedBy string) EventView{
	return EventView{
			Id:event.Id,
			EventName:event.EventName,
			BookedBy:bookedBy}
}

func eventIsBooked(db* gorm.DB, event Event) (bool, Booking){
	var v = findBookForEventId(db, event.Id)
	if v.Id != 0{
		return true, v
	}

	return false, v
}