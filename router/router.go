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

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	// User
	user := api.Group("/user")
	user.Get("/:id", handler.GetUser)
	user.Post("/", handler.CreateUser)
	user.Patch("/:id", middleware.Protected(), handler.UpdateUser)
	user.Delete("/:id", middleware.Protected(), handler.DeleteUser)

	// Product
	product := api.Group("/product")
	product.Get("/", handler.GetAllProjects)
	product.Get("/:id", handler.GetProject)
	product.Post("/", handler.CreateProject)
	product.Delete("/:id", middleware.Protected(), handler.DeleteProject)
}
