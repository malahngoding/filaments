package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/malahngoding/filaments/config"
)

// Createtopic api handler
func Createtopic(c *fiber.Ctx) error {
	client := hedera.ClientForTestnet()
	client.SetOperator(config.HederaID(), config.HederaPrivateKey())

	transactionResponse, err := hedera.NewTopicCreateTransaction().
		Execute(client)
	if err != nil {
		println(err.Error(), ": error creating topic")
		return err
	}

	transactionReceipt, err := transactionResponse.GetReceipt(client)
	if err != nil {
		println(err.Error(), ": error getting topic create receipt")
		return err
	}

	topicID := *transactionReceipt.TopicID

	return c.JSON(fiber.Map{
		"messages": "Hello Future",
		"payload":  fiber.Map{"topicId": topicID.String()},
		"status":   "OK",
	})
}

// SubmitMessageToTopic api handler
func SubmitMessageToTopic(c *fiber.Ctx) error {
	client := hedera.ClientForTestnet()
	client.SetOperator(config.HederaID(), config.HederaPrivateKey())

	type ConsensusMessage struct {
		TopicId string `json:"topicId" xml:"topicId" form:"topicId"`
		Message string `json:"message" xml:"message" form:"message"`
		Memo    string `json:"memo" xml:"memo" form:"memo"`
	}

	sc := new(ConsensusMessage)
	if err := c.BodyParser(sc); err != nil {
		return err
	}

	hashTopicId, err := hedera.TopicIDFromString(sc.TopicId)
	if err != nil {
		return err
	}

	submitMessage, err := hedera.NewTopicMessageSubmitTransaction().
		SetMessage([]byte(sc.Message)).
		SetTopicID(hashTopicId).
		SetTransactionMemo(sc.Memo).
		Execute(client)
	if err != nil {
		println(err.Error(), ": error submitting to topic")
		return err
	}

	//Get the receipt of the transaction
	receipt, err := submitMessage.GetReceipt(client)
	if err != nil {
		println(err.Error(), ": error submitting to topic")
		return err
	}

	transactionStatus := receipt.Status

	return c.JSON(fiber.Map{
		"messages": "Hello Future",
		"payload": fiber.Map{
			"messages": "The message transaction status " + transactionStatus.String(),
		},
		"status": "OK",
	})
}
