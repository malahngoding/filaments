package main

import (
	"fmt"

	"github.com/malahngoding/filaments/server"
)

func main() {
	fmt.Println("Run Filaments")
	go server.ServerWS()
	server.ServerHTTP()
}
