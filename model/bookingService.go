package model

import (
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
	"encoding/json"
	"net/http"
	"github.com/go-martini/martini"
	"fmt"
)

func BookEvent(request* http.Request, db *gorm.DB, render render.Render){
	decoder := json.NewDecoder(request.Body)
	defer request.Body.Close()

	var booking Booking
	err := decoder.Decode(&booking)

	if err != nil {
		render.Text(http.StatusBadRequest, "Bad JSON encoding")
		return
	}

	if booking.BookedBy == ""{
		render.Text(http.StatusBadRequest, "BookedBy cannot be empty")
		return
	}

	if findEventById(db, fmt.Sprint(booking.EventId)).Id == 0{
		render.Text(http.StatusNotFound, "Event not found")
		return
	}

	if findBookForEventId(db, fmt.Sprint(booking.EventId)).Id != 0{
		render.Text(http.StatusBadRequest, "Event already booked")
		return
	}

	bookEvent(db, booking)

	render.Text(http.StatusCreated, "book created")
}


func UnBookEvent(params martini.Params, db *gorm.DB, render render.Render){
	var booking Booking

	if findBookForEventId(db, params["id"]).Id == 0{
		render.Text(http.StatusNotFound, "Event is not booked!")
		return
	}

	unBookEvent(db, booking)

	render.Text(http.StatusOK, "Booking deleted")
}