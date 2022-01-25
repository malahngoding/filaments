package config

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

func HederaID() hedera.AccountID {
	operatorID, err := hedera.AccountIDFromString(Config("HEDERA_ACCOUNT_ID"))
	if err != nil {
		fmt.Println(operatorID)
		panic(err)
	}
	return operatorID
}

func HederaPublicKey() hedera.PublicKey {
	operatorPublicKey, err := hedera.PublicKeyFromString(Config("HEDERA_ACCOUNT_PUBLIC_KEY"))
	if err != nil {
		fmt.Println(operatorPublicKey)
		panic(err)
	}
	return operatorPublicKey
}

func HederaPrivateKey() hedera.PrivateKey {
	operatorPrivateKey, err := hedera.PrivateKeyFromString(Config("HEDERA_ACCOUNT_PRIVATE_KEY"))
	if err != nil {
		fmt.Println(operatorPrivateKey)
		panic(err)
	}
	return operatorPrivateKey
}
