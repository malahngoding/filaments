package server

import (
	"fmt"
	"log"

	"github.com/malahngoding/filaments/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func ServerHTTP() {
	fmt.Println("Running HTTP on :4444")
	app := fiber.New()
	app.Use(cors.New())
	app.Use(limiter.New())

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":4444"))
}
