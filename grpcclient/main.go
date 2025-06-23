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

	response, err := userClient.CreateUser(context.Background(), &user.User{
		Id:      1, //! input yang berada di user.proto
		Age:     13,
		Balance: 130000,
		Address: &user.Address{
			Id:          123,
			FullAddress: "Jln. Surabaya",
			Province:    "Jawa Timur",
			City:        "Surabaya",
		},
	}) //!CreateUser: method yang ada di proto (hasil generate)
	if err != nil {
		log.Fatal("Error calling user client ", err)
	}
	log.Println("Got message from server: ", response.Message)

}
