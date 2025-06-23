package main

import (
	"context"
	"grpc-course-protobuf/pb/chat"
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

	stream, err := chatClient.SendMessage(context.Background())
	if err != nil {
		log.Fatal("Failed to send message", err)
	}

	err = stream.Send(&chat.ChatMessage{ //!mengirim pesan 1
		UserId:  123, //!kita bisa mengirim data chat
		Content: "Hello from client, pesan pertama",
	})
	if err != nil {
		log.Fatal("Failed to send via stream", err)
	}

	err = stream.Send(&chat.ChatMessage{ //!mengirim pesan 2
		UserId:  123, //!kita bisa mengirim data chat
		Content: "Hello again, pesan kedua",
	})
	if err != nil {
		log.Fatal("Failed to send via stream", err)
	}

	res, err := stream.CloseAndRecv() //! mengirim info bahwa stream telah selesai dan ditutup
	if err != nil {
		log.Fatal("Failed close", err)
	}
	log.Println("Connection is closed. Message: ", res.Message)

}
