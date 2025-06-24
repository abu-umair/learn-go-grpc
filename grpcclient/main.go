package main

import (
	"context"
	"errors"
	"grpc-course-protobuf/pb/chat"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	clientConn, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials())) //insecure.NewCredentials(): jika dilingkungan develop
	if err != nil {
		log.Fatal("Failed create client", err)
	}

	chatClient := chat.NewChatServiceClient(clientConn)

	stream, err := chatClient.ReceiveMessage(context.Background(), &chat.ReceiveMessageRequest{
		UserId: 30,
	})
	if err != nil {
		log.Fatal("Failed to send message", err)
	}

	for { //!perulangan untuk menerima banyaknya pesan dari server
		msg, err := stream.Recv() //? menerima data" yang dikirim dari server

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatal("Failed to receive message ", err)
		}

		log.Printf("Got message to %d content %s", msg.UserId, msg.Content)
	}

}
