package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	clientConn, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials())) //insecure.NewCredentials(): jika dilingkungan develop
	if err != nil{
		log.Fatal("Failed create client", err)
	}
}
