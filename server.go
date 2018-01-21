package main

import (
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/cors"
	"booking-go-rest-api/model"
	"github.com/go-martini/martini"
)

func main() {
	var db = model.InitDb()
	defer db.Close()

	server := martini.Classic()
	server.Map(db)
	server.Use(render.Renderer())
	server.Use(cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "DELETE", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	server.Group("/", func (r martini.Router) {
		r.Get("event/", model.GetAllEvents)
		r.Delete("event/:id", model.DeleteEvent)
		r.Put("event/", model.CreateEvent)
		r.Put("bookEvent/", model.BookEvent)
		r.Delete("unBookEvent/:id", model.UnBookEvent)
	})

	server.Run()
}
