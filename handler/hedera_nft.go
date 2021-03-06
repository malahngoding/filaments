package handler

import (
	"github.com/malahngoding/filaments/config"

	"github.com/gofiber/fiber/v2"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

// HelloNFT api handler
func HelloNFT(c *fiber.Ctx) error {
	client := hedera.ClientForTestnet()
	client.SetOperator(config.HederaID(), config.HederaPrivateKey())

	return c.JSON(fiber.Map{
		"messages": "Hello NFT",
		"payload": fiber.Map{
			"newToken": fiber.Map{
				"name": "b5ef6c8b39db0cd25f6f683a1425ec6f Token",
			}},
		"status": "OK",
	})
}

// CreateNFT api handler
func CreateNFT(c *fiber.Ctx) error {
	client := hedera.ClientForTestnet()
	client.SetOperator(config.HederaID(), config.HederaPrivateKey())

	return c.JSON(fiber.Map{
		"messages": "NFT created successfully",
		"payload": fiber.Map{
			"newToken": fiber.Map{
				"name": "b5ef6c8b39db0cd25f6f683a1425ec6f Token",
			}},
		"status": "OK",
	})
}
