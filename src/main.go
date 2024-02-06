package main

import (
	"go-fiber-restapi/src/books"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/book", books.Details)
	app.Get("/books", books.Index)
	app.Post("/book", books.Create)
	app.Put("/book", books.Update)
	app.Delete("/book", books.Delete)

	app.Listen(":3000")
}