package handler

import (
	"filaments/config"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

// HelloAccount api handler
func HelloAccount(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"messages": "Hello Account",
		"payload":  fiber.Map{"empty": true},
		"status":   "OK",
	})
}

// HelloAccount api handler
func CreateAccount(c *fiber.Ctx) error {
	client := hedera.ClientForTestnet()
	client.SetOperator(config.HederaID(), config.HederaPrivateKey())

	mnemonic, err := hedera.GenerateMnemonic24()
	if err != nil {
		panic(err)
	}

	privateKey, err := hedera.PrivateKeyFromMnemonic(mnemonic, "")
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.PublicKey()

	//Create an account with 1,000 hbar
	AccountCreateTransaction := hedera.NewAccountCreateTransaction().
		SetKey(publicKey).SetInitialBalance(hedera.NewHbar(1000))

	//Return the key on the account
	accountKey, err := AccountCreateTransaction.GetKey()
	if err != nil {
		panic(err)
	}

	return c.JSON(fiber.Map{
		"messages": fmt.Sprintln(accountKey),
		"payload": fiber.Map{
			"newAccount": fiber.Map{
				"id": accountKey,
			}},
		"status": "OK",
	})
}
