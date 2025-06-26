package main

import (
	"context"
	"grpc-course-protobuf/pb/chat"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	clientConn, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials())) //insecure.NewCredentials(): jika dilingkungan develop
	if err != nil {
		log.Fatal("Failed create client", err)
	}

	chatClient := chat.NewChatServiceClient(clientConn)

	stream, err := chatClient.Chat(context.Background())
	if err != nil {
		log.Fatal("Failed to send message", err)
	}

	err = stream.Send(&chat.ChatMessage{
		UserId:  123,
		Content: "Hello this is client",
	})
	if err != nil {
		log.Fatalf("Failed to send message %v", err)
	}

	msg, err := stream.Recv()
	if err != nil {
		log.Fatalf("Failed to receive message %v", err)
	}

	log.Printf("Got reply from server %d content %s\n", msg.UserId, msg.Content)

	time.Sleep(5 * time.Second)

	err = stream.Send(&chat.ChatMessage{
		UserId:  123,
		Content: "Hello this is client again",
	})
	if err != nil {
		log.Fatalf("Failed to send message %v", err)
	}

	msg, err = stream.Recv()
	if err != nil {
		log.Fatalf("Failed to receive message %v", err)
	}

	log.Printf("Got reply from server %d content %s\n", msg.UserId, msg.Content)

}
