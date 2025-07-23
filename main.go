package main

import (
	"context"
	"errors"
	"fmt"
	"grpc-course-protobuf/pb/chat"
	"grpc-course-protobuf/pb/common"
	"grpc-course-protobuf/pb/user"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type userService struct {
	user.UnimplementedUserServiceServer //mengiinitialkan semua API USER (mungkin seperti resource di route laravel)
}

func (us *userService) CreateUser(ctx context.Context, userRequest *user.User) (*user.CreateResponse, error) {
	if userRequest.Age < 1 { //?membuat validasi age
		return &user.CreateResponse{
			Base: &common.BaseResponse{
				StatusCode: 400,
				IsSuccess:  false,
				Message:    "Validation Error",
			},
		}, nil
	}

	// return nil, status.Errorf(codes.Internal, "Server is bugged") //?membuat example error internal

	log.Println("CreateUser is running")
	return &user.CreateResponse{
		Base: &common.BaseResponse{
			StatusCode: 0,
			IsSuccess:  true,
			Message:    "Success Create User",
		},
	}, nil
}

type chatService struct {
	chat.UnimplementedChatServiceServer
}

func (cs *chatService) SendMessage(stream grpc.ClientStreamingServer[chat.ChatMessage, chat.ChatResponse]) error { //?mengambil template di chat.grpc.pb.go

	for { //!perulangan untuk menerima banyaknya pesan yang dikirim
		req, err := stream.Recv() //? singkatan dari recived (menerima pesan)
		if err != nil {
			if errors.Is(err, io.EOF) { //!jika pesan sudah selesai (menerima komunikasi mengakhiri dari client)
				break //?io.EOF: ada di dokumentasi
				//!ketika mode prod, lebih baik menambahkan time out beberapa detik (misal HP nya ngehang, jadi tidak ada komunikasi)
			}
			return status.Errorf(codes.Unknown, "Error receiving message %v", err)
		}
		log.Printf("Receive message: %s, to %d", req.Content, req.UserId)
	}

	return stream.SendAndClose(&chat.ChatResponse{
		Message: "Thanks for the messages!",
	})
}

// !method utk server streaming (server mengirim banyak data dalam satu kali koneksi).
// ? terdapat 2 objek: req dan stream
func (cs *chatService) ReceiveMessage(req *chat.ReceiveMessageRequest, stream grpc.ServerStreamingServer[chat.ChatMessage]) error { //?mengambil template di chat.grpc.pb.go

	log.Printf("Got connection request from %d\n", req.UserId) //?UserId: id user yang mengirim request (dpt dilihat di chat.proto)

	for i := 0; i < 10; i++ {
		err := stream.Send(&chat.ChatMessage{ //?Send: untuk mengirim data ke client, kemudian mengakses objek stream
			UserId:  123,
			Content: "Hi, from server akan mengrim sebanyak 10 kali pesan", //?message ke 1
		})
		if err != nil {
			return status.Errorf(codes.Unknown, "error sending message to client %v", err)
		}
	}

	return nil
}

func (cs *chatService) Chat(stream grpc.BidiStreamingServer[chat.ChatMessage, chat.ChatMessage]) error { //?mengambil template di chat.grpc.pb.go
	for {
		msg, err := stream.Recv() //?object stream jg ada (diambil dari chat.grpc.pb.go)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return status.Error(codes.Unknown, fmt.Sprintf("error receiving message: %v", err))
		}

		log.Printf("Got message from %d content: %s", msg.UserId, msg.Content)

		time.Sleep(2 * time.Second)

		err = stream.Send(&chat.ChatMessage{
			UserId:  50,
			Content: "Reply from server",
		})
		if err != nil {
			return status.Error(codes.Unknown, fmt.Sprintf("error sending message: %v", err))
		}

		err = stream.Send(&chat.ChatMessage{
			UserId:  50,
			Content: "Reply from server #2",
		})
		if err != nil {
			return status.Error(codes.Unknown, fmt.Sprintf("error sending message: %v", err))
		}
	}

	return nil
}

func main() {
	// lis, err := net.Listen(network: "tcp", address: ":8080") //ditutorial seperti ini
	lis, err := net.Listen("tcp", ":8081") //harus sama dengan port di grpcclient/main.go
	if err != nil {
		log.Fatal("There is error in your net listen ", err)
	}

	serv := grpc.NewServer()

	user.RegisterUserServiceServer(serv, &userService{})

	chat.RegisterChatServiceServer(serv, &chatService{})

	reflection.Register(serv)

	if err := serv.Serve(lis); err != nil {
		// log.Fatav...: "Error running server ", err)
		log.Fatal("Error running server ", err)
	}
}
