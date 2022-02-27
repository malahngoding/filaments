package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/malahngoding/filaments/config"
)

func ServerSymmetry() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.GetReqHeaders()
		if validateSymmetry(token["Authorization"]) {
			return c.Next()
		}
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"status":  "error",
				"message": "Failed Authentication",
				"data":    nil,
			})
	}
}

func validateSymmetry(token string) bool {
	trimmed := strings.TrimPrefix(token, "Bearer instead_")
	return trimmed == config.InsteadToken()
}
