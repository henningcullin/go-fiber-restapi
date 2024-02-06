package main

import (
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/hello", func(c fiber.Ctx) error {
		return c.SendString("Hello, World from Fiber!")
	})

	app.Listen(":3000")
}