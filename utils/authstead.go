package utils

import (
	"context"

	"github.com/malahngoding/filaments/config"
	"google.golang.org/grpc/metadata"
)

func Authstead(ctx context.Context) bool {
	var values []string
	var token string
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		values = md.Get("authorization")
	}

	if len(values) > 0 {
		token = values[0]
	}

	if token == config.InsteadToken() {
		return true
	} else {
		return false
	}
}
