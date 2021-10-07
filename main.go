package main

import "fmt"

func main() {
	fmt.Println("Run Filaments")
	go ServerRPC()
	ServerHTTP()
}
