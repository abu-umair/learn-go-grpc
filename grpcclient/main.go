package main

import (
	"context"
	"grpc-course-protobuf/pb/user"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	clientConn, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials())) //insecure.NewCredentials(): jika dilingkungan develop
	if err != nil {
		log.Fatal("Failed create client", err)
	}

	userClient := user.NewUserServiceClient(clientConn)

	res, err := userClient.CreateUser(context.Background(), &user.User{
		Age: -1, //?inputnya negatif, jadi error
	})
	if err != nil {
		log.Println("Failed to send message", err)//?menganggap responsenya error
		log.Println(res)//? nilainya nil / gak ada response
		return
	}

	log.Println("Response from server ", res.Message) //?response dari server

}
