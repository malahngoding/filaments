package middleware

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/malahngoding/filaments/config"
	"github.com/malahngoding/filaments/utils"
)

var ctx = context.Background()

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
	opt, _ := redis.ParseURL(config.RedisAPI())
	client := redis.NewClient(opt)
	val := client.Get(ctx, token).Val()
	if val != "GRANT" {
		return false
	}
	key := []byte(config.InsteadToken())

	trimmed := strings.TrimPrefix(token, "Bearer instead_")
	value, err := utils.Decrypt(key, trimmed)
	if err != nil {
		fmt.Println(err)
	}
	if strings.Contains(value, "GITHUB") {
		return true
	}
	if strings.Contains(value, "GITHUB") {
		return true
	}
	return false
}
