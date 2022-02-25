package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Authenticated() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.GetReqHeaders()
		if validateToken(token["Authorization"]) {
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

func validateToken(token string) bool {
	trimmed := strings.TrimPrefix(token, "Bearer instead_")
	return trimmed == "U2FsdGVkX18nG9YH1sTHwI9lPqMAeT1q4B8naEcJ2zTakhlptqaczLHV5DFHVOwU"
}
