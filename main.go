package main

import (
	"filaments/server"
	"fmt"
)

func main() {
	fmt.Println("Run Filaments")
	go server.ServerRPC()
	go server.ServerWS()
	server.ServerHTTP()
}
