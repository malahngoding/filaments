package main

import (
	"filaments/router"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func ServerHTTP() {
	fmt.Println("Running HTTP on :4444")
	app := fiber.New()
	app.Use(cors.New())

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":4444"))
}
