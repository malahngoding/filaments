package server

import (
	"fmt"
	"log"
	"time"

	"github.com/malahngoding/filaments/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func ServerHTTP() {
	fmt.Println("Running HTTP on :4444")
	app := fiber.New()
	app.Use(cors.New())
	app.Use(limiter.New(limiter.Config{
		Max:        120,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusTooManyRequests)
		},
		SkipFailedRequests:     false,
		SkipSuccessfulRequests: false,
		LimiterMiddleware:      limiter.SlidingWindow{},
	}))

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":4444"))
}
