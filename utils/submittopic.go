package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/malahngoding/filaments/config"
)

func handleID(id int) string {
	switch id {
	case 1:
		return "LOGGING_IN"
	}
	return "ERROR"
}

func SubmitTopic(id int, message string) {
	topicID := "0.0.34291358"
	client := hedera.ClientForTestnet()
	client.SetOperator(config.HederaID(), config.HederaPrivateKey())

	t := time.Now().Local().Format(time.RFC850)
	s := []string{t, "@", handleID(id), "-", message}
	printable := strings.Join(s, " ")

	hashTopicId, err := hedera.TopicIDFromString(topicID)
	if err != nil {
		fmt.Println(err)
	}

	submitMessage, err := hedera.NewTopicMessageSubmitTransaction().
		SetMessage([]byte(printable)).
		SetTopicID(hashTopicId).
		SetTransactionMemo(printable).
		Execute(client)
	if err != nil {
		fmt.Println(err.Error(), ": error submitting to topic")
	}

	receipt, err := submitMessage.GetReceipt(client)
	if err != nil {
		fmt.Println(err.Error(), ": error submitting to topic")
	}

	transactionStatus := receipt.Status
	fmt.Println(transactionStatus.String())
}
