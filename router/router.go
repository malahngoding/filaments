package router

import (
	"github.com/malahngoding/filaments/handler"
	"github.com/malahngoding/filaments/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	// Handshake
	handshake := api.Group("/handshake")
	handshake.Post("/", middleware.ServerSymmetry(), handler.Handshake)

	// Hedera Endpoint
	hedera := api.Group("/hedera")
	// Hedera Account
	hedera.Post("/account/create", middleware.ServerSymmetry(), handler.CreateAccount)
	hedera.Get("/account/:id", middleware.Authenticated(), handler.HelloAccount)
	// Hedera Topics
	hedera.Post("/topic/create", middleware.ServerSymmetry(), handler.Createtopic)
	hedera.Post("/topic/submit", middleware.Authenticated(), handler.SubmitMessageToTopic)
	// Hedera Token
	hedera.Post("/token/create", middleware.ServerSymmetry(), handler.CreateToken)
	hedera.Post("/token/associate", middleware.Authenticated(), handler.AssociateToken)
	hedera.Post("/token/disassociate", middleware.Authenticated(), handler.DisassociateToken)
	hedera.Get("/token/", middleware.Authenticated(), handler.HelloToken)
	// Hedera NFT
	hedera.Post("/nft/create", middleware.ServerSymmetry(), handler.CreateToken)
	hedera.Get("/nft", middleware.Authenticated(), handler.HelloToken)
	// Hedera Smart Contracts
	hedera.Post("/smartcontracts/call", middleware.ServerSymmetry(), handler.ContractCall)
	hedera.Get("/smartcontracts", middleware.ServerSymmetry(), handler.HelloSmartContracts)

	// Comments End Point
	comment := api.Group(("/comments"))
	comment.Get("/", middleware.ServerSymmetry(), handler.GetComments)
	comment.Post("/add", middleware.Authenticated(), handler.AddComment)

	// Quiz Box Test with SQLite
	quizBox := api.Group("/quiz-box")
	quizBox.Get("/crud", handler.GetCrud)
	quizBox.Post("/crud", handler.PostCrud)
	quizBox.Patch("/crud", handler.PatchCrud)
	quizBox.Delete("/crud", handler.DeleteCrud)
}
