package handler

import (
	"fmt"

	"github.com/malahngoding/filaments/config"

	"github.com/gofiber/fiber/v2"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

// HelloAccount api handler
func HelloAccount(c *fiber.Ctx) error {
	client := hedera.ClientForTestnet()
	client.SetOperator(config.HederaID(), config.HederaPrivateKey())

	accountID, err := hedera.AccountIDFromString(c.Params("id"))
	if err != nil {
		panic(err)
	}

	query := hedera.NewAccountBalanceQuery().SetAccountID(accountID)

	accountBalance, err := query.Execute(client)
	if err != nil {
		panic(err)
	}

	return c.JSON(fiber.Map{
		"messages": fmt.Sprintf("Hello Account %s", accountID),
		"payload": fiber.Map{
			"acccount": fiber.Map{
				"accountID": fmt.Sprintf("%d.%d.%d", accountID.Shard, accountID.Realm, accountID.Account),
				"balance":   accountBalance.Hbars.String(),
			},
		},
		"status": "OK",
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

	AccountCreateTransaction := hedera.NewAccountCreateTransaction().
		SetKey(publicKey).SetInitialBalance(hedera.NewHbar(0))

	txResponse, err := AccountCreateTransaction.Execute(client)
	if err != nil {
		panic(err)
	}

	receipt, err := txResponse.GetReceipt(client)
	if err != nil {
		panic(err)
	}

	newAccountId := *receipt.AccountID

	if err != nil {
		panic(err)
	}

	return c.JSON(fiber.Map{
		"messages": "Account created successfully",
		"payload": fiber.Map{
			"newAccount": fiber.Map{
				"id":         newAccountId,
				"mnemonic":   mnemonic.Words(),
				"privateKey": privateKey.String(),
			}},
		"status": "OK",
	})
}
