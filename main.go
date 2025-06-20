package main

import (
	"log"
	"net"
)

func main() {
	// lis, err := net.Listen(network: "tcp", address: ":8080") //ditutorial seperti ini
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("There is error in your net listen ", err)
	}
}
