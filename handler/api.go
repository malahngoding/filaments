package handler

import (
	"filaments/utils"

	"github.com/gofiber/fiber/v2"
)

// Hello hanlde api status
func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Hello i'm ok!",
		"num":     utils.Add(12, 50),
	})
}
