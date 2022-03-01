package handler

import (
	"github.com/gofiber/fiber/v2"
)

// Hello api handle
func Hello(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"messages": "Hello Future",
		"payload":  fiber.Map{"empty": nil},
		"status":   "OK",
	})
}
