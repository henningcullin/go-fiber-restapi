package books

import (
	"github.com/gofiber/fiber/v3"
)

func Details(c fiber.Ctx) error {
	return c.SendString("Book details")
}

func Index(c fiber.Ctx) error {
	return c.SendString("All books")
}

func Create(c fiber.Ctx) error {
	return c.SendString("Book created")
}

func Update(c fiber.Ctx) error {
	return c.SendString("Book updated")
}

func Delete(c fiber.Ctx) error {
	return c.SendString("Book deleted")	
}