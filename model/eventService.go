package model

import (
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
	"net/http"
	"github.com/go-martini/martini"
	"encoding/json"
	"unicode/utf8"
	"strconv"
)

func GetAllEvents(db* gorm.DB, render render.Render){
	var eventView = getAllEventsView(db)

	if eventView == nil{
		render.JSON(http.StatusOK, make([]Event, 0))
		return
	}

	render.JSON(http.StatusOK, eventView)
}

func DeleteEvent(params martini.Params, db* gorm.DB, render render.Render){
	var event Event
	eventId, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		render.Text(http.StatusBadRequest, "Bad JSON encoding")
		return
	}

	if findEventById(db, eventId).Id == 0 {
		render.Text(http.StatusNotFound, "Event not found")
		return
	}

	if findBookForEventId(db, eventId).Id != 0{
		render.Text(http.StatusNotFound, "Event cannot be deleted, it's booked")
		return
	}

	event.Id = eventId
	deleteEvent(db, &event)

	render.Text(http.StatusOK, "Event deleted")
}

func CreateEvent(request* http.Request, db* gorm.DB, render render.Render){
	decoder := json.NewDecoder(request.Body)
	defer request.Body.Close()

	var event Event
	err := decoder.Decode(&event)

	if err != nil {
		render.Text(http.StatusBadRequest, "Bad JSON encoding")
		return
	}

	if findEventByName(db, event.EventName).Id != 0{
		render.Text(http.StatusOK, "Event already exists")
		return
	}

	if event.EventName == ""{
		render.Text(http.StatusOK, "Event name cannot be empty")
		return
	}

	if utf8.RuneCountInString(event.EventName) > 100{
		render.Text(http.StatusOK, "Event is too length, max size is 100")
		return
	}

	createEvent(db, &event)

	render.JSON(http.StatusCreated, event)
}