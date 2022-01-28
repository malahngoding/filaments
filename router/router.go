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

	// Hedera Endpoint
	hedera := api.Group("/hedera")
	// Hedera Account
	hedera.Post("/account/create", handler.CreateAccount)
	hedera.Get("/account/:id", handler.HelloAccount)
	// Hedera Token
	hedera.Post("/token/create", handler.CreateToken)
	hedera.Get("/token", handler.HelloToken)
	// Hedera NFT
	hedera.Post("/nft/create", handler.CreateToken)
	hedera.Get("/nft", handler.HelloToken)

}
