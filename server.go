package main

import (
	"github.com/martini-contrib/render"
	"booking-go-rest-api/model"
	"github.com/go-martini/martini"
)

func main() {
	db := model.InitDb()
	defer db.Close()

	server := martini.Classic()
	server.Map(db)
	server.Use(render.Renderer())

	server.Group("/", func (r martini.Router) {
		r.Get("events/", model.GetAllEvents)
		r.Delete("events/:id", model.DeleteEvent)
		r.Put("events/", model.CreateEvent)
		r.Put("bookEvent/", model.BookEvent)
		r.Delete("unBookEvent/:id", model.UnBookEvent)
	})

	server.Run()
}
