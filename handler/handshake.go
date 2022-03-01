package handler

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/malahngoding/filaments/config"
	"github.com/malahngoding/filaments/utils"
)

type HandhsakeRequest struct {
	Identification string `json:"identification" xml:"identification" form:"identification"`
	Provider       string `json:"provider" xml:"provider" form:"provider"`
}

var ctx = context.Background()

// Hello api handle
func Handshake(c *fiber.Ctx) error {
	hr := new(HandhsakeRequest)
	if err := c.BodyParser(hr); err != nil {
		return err
	}
	s := []string{hr.Identification, "_", "_", hr.Provider}
	combined := strings.Join(s, "")

	key := []byte(config.InsteadToken())
	encrypted, err := utils.Encrypt(key, combined)
	if err != nil {
		fmt.Println(err)
	}
	opt, _ := redis.ParseURL(config.RedisAPI())
	client := redis.NewClient(opt)
	client.Set(ctx, encrypted, "GRANT", 0)

	return c.JSON(fiber.Map{
		"messages": "Hello Future",
		"payload": fiber.Map{
			"token": encrypted,
		},
		"status": "OK",
	})
}
