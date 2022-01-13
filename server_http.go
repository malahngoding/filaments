package main

import (
	"filaments/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func ServerHTTP() {
	app := fiber.New()
	app.Use(cors.New())

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":4444"))

}
