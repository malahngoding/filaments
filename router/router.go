package router

import (
	"filaments/handler"
	"filaments/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	// Project
	project := api.Group("/project")
	project.Get("/", handler.GetAllProjects)
	project.Get("/:id", handler.GetProject)
	project.Post("/", handler.CreateProject)
	project.Delete("/:id", middleware.Protected(), handler.DeleteProject)

	// Hedera Account
	hedera := api.Group("/hedera")
	hedera.Get("/account", handler.HelloAccount)
	hedera.Get("/account/create", handler.CreateAccount)

}
