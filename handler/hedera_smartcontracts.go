package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/malahngoding/filaments/config"
)

// ContractCall api handler
func ContractCall(c *fiber.Ctx) error {
	client := hedera.ClientForTestnet()
	client.SetOperator(config.HederaID(), config.HederaPrivateKey())

	type SmartContracts struct {
		SolidityAddress string `json:"solidityAddress" xml:"solidityAddress" form:"solidityAddress"`
		Action          string `json:"action" xml:"action" form:"action"`
	}

	sc := new(SmartContracts)
	if err := c.BodyParser(sc); err != nil {
		return err
	}
	fmt.Println(sc.SolidityAddress)
	newContractID, err := hedera.ContractIDFromSolidityAddress(sc.SolidityAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println(newContractID)
	contractQuery, err := hedera.NewContractCallQuery().
		SetContractID(newContractID).
		SetGas(100000).
		SetQueryPayment(hedera.NewHbar(1)).
		SetFunction(sc.Action, nil).
		Execute(client)

	if err != nil {
		println(err.Error(), ": error executing contract call query")
		return err
	}

	// Get a string from the result at index 0
	getMessage := contractQuery.GetString(0)

	return c.JSON(fiber.Map{
		"messages": "Hello Smart Contracts Call",
		"payload": fiber.Map{
			"The contract message: ": getMessage,
		},
		"status": "OK",
	})
}

// ContractExecute api handler
func ContractExecute(c *fiber.Ctx) error {
	client := hedera.ClientForTestnet()
	client.SetOperator(config.HederaID(), config.HederaPrivateKey())

	type SmartContracts struct {
		SolidityAddress string `json:"solidityAddress" xml:"solidityAddress" form:"solidityAddress"`
		Action          string `json:"action" xml:"action" form:"action"`
	}

	sc := new(SmartContracts)
	if err := c.BodyParser(sc); err != nil {
		return err
	}
	fmt.Println(sc.SolidityAddress)
	newContractID, err := hedera.ContractIDFromSolidityAddress(sc.SolidityAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println(newContractID)
	contractQuery, err := hedera.NewContractCallQuery().
		SetContractID(newContractID).
		SetGas(100000).
		SetQueryPayment(hedera.NewHbar(1)).
		SetFunction(sc.Action, nil).
		Execute(client)

	if err != nil {
		println(err.Error(), ": error executing contract call query")
		return err
	}

	// Get a string from the result at index 0
	getMessage := contractQuery.GetString(0)

	return c.JSON(fiber.Map{
		"messages": "Hello Smart Contracts Execute",
		"payload": fiber.Map{
			"The contract message: ": getMessage,
		},
		"status": "OK",
	})
}

// HelloSmartContracts api handler
func HelloSmartContracts(c *fiber.Ctx) error {
	client := hedera.ClientForTestnet()
	client.SetOperator(config.HederaID(), config.HederaPrivateKey())

	return c.JSON(fiber.Map{
		"messages": "Hello Smart Contracts",
		"payload": fiber.Map{
			"smartcontracts": true,
		},
		"status": "OK",
	})
}
