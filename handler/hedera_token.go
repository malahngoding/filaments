package handler

import (
	"fmt"

	"github.com/malahngoding/filaments/config"

	"github.com/gofiber/fiber/v2"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

// HelloToken api handler
func HelloToken(c *fiber.Ctx) error {
	client := hedera.ClientForTestnet()
	client.SetOperator(config.HederaID(), config.HederaPrivateKey())

	return c.JSON(fiber.Map{
		"messages": "Hello Token",
		"payload": fiber.Map{
			"newToken": fiber.Map{
				"name": "Malah Ngoding Token",
			}},
		"status": "OK",
	})
}

// CreateToken api handler
func CreateToken(c *fiber.Ctx) error {
	type Treasury struct {
		TreasuryId  string `json:"treasuryId" xml:"treasuryId" form:"treasuryId"`
		TreasuryKey string `json:"treasuryKey" xml:"treasuryKey" form:"treasuryKey"`
	}

	t := new(Treasury)
	if err := c.BodyParser(t); err != nil {
		return err
	}

	client := hedera.ClientForTestnet()
	client.SetOperator(config.HederaID(), config.HederaPrivateKey())

	treasuryAccountId, err := hedera.AccountIDFromString(t.TreasuryId)
	if err != nil {
		panic(err)
	}

	treasuryKey, err := hedera.PrivateKeyFromString(t.TreasuryKey)
	if err != nil {
		panic(err)
	}

	mnemonic, err := hedera.GenerateMnemonic24()
	if err != nil {
		panic(err)
	}

	supplyKey, err := hedera.PrivateKeyFromMnemonic(mnemonic, "")
	if err != nil {
		panic(err)
	}

	tokenCreateTx, err := hedera.NewTokenCreateTransaction().
		SetTokenName("Malah Ngoding Token").
		SetTokenSymbol("MNT").
		SetTokenType(hedera.TokenTypeFungibleCommon).
		SetDecimals(6).
		SetInitialSupply(50000000000000).
		SetTreasuryAccountID(treasuryAccountId).
		SetSupplyType(hedera.TokenSupplyTypeInfinite).
		SetSupplyKey(supplyKey).
		SetTokenMemo("Malah Ngoding Token for Malah Ngoding Instead").
		FreezeWith(client)
	if err != nil {
		panic(err)
	}

	tokenCreateSign := tokenCreateTx.Sign(treasuryKey)

	tokenCreateSubmit, err := tokenCreateSign.Execute(client)
	if err != nil {
		panic(err)
	}

	tokenCreateRx, err := tokenCreateSubmit.GetReceipt(client)
	if err != nil {
		panic(err)
	}

	tokenId := *tokenCreateRx.TokenID

	return c.JSON(fiber.Map{
		"messages": "Token created successfully",
		"payload": fiber.Map{
			"newToken": fiber.Map{
				"name": "Malah Ngoding Token",
				"treasuryId": fmt.Sprintf(
					"%d.%d.%d",
					treasuryAccountId.Shard,
					treasuryAccountId.Realm,
					treasuryAccountId.Account,
				),
				"treasuryKey": treasuryKey.String(),
				"supplyKey":   supplyKey.String(),
				"tokenId": fmt.Sprintf(
					"%d.%d.%d",
					tokenId.Shard,
					tokenId.Realm,
					tokenId.Token,
				),
			}},
		"status": "OK",
	})
}

// CreateToken api handler
func AssociateToken(c *fiber.Ctx) error {
	type Association struct {
		TargetId  string `json:"targetId" xml:"targetId" form:"targetId"`
		TargetKey string `json:"targetKey" xml:"targetKey" form:"targetKey"`
		TokenId   string `json:"tokenId" xml:"tokenId" form:"tokenId"`
	}

	a := new(Association)
	if err := c.BodyParser(a); err != nil {
		return err
	}

	client := hedera.ClientForTestnet()
	client.SetOperator(config.HederaID(), config.HederaPrivateKey())

	targetId, err := hedera.AccountIDFromString(a.TargetId)
	if err != nil {
		panic(err)
	}

	targetKey, err := hedera.PrivateKeyFromString(a.TargetKey)
	if err != nil {
		panic(err)
	}

	tokenId, err := hedera.TokenIDFromString(a.TokenId)
	if err != nil {
		panic(err)
	}

	associateMNT, err := hedera.NewTokenAssociateTransaction().
		SetAccountID(targetId).
		SetTokenIDs(tokenId).
		FreezeWith(client)
	if err != nil {
		panic(err)
	}

	signTx := associateMNT.Sign(targetKey)

	associateMNTSubmit, err := signTx.Execute(client)
	if err != nil {
		panic(err)
	}

	associateMNTrx, err := associateMNTSubmit.GetReceipt(client)
	if err != nil {
		panic(err)
	}

	return c.JSON(fiber.Map{
		"messages": "Hello Association",
		"payload": fiber.Map{
			"newToken": fiber.Map{
				"name": fmt.Sprintf("MNT Associated to account %s", targetId.String()),
				"id":   associateMNTrx.Status,
			}},
		"status": "OK",
	})
}
