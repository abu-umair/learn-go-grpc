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
	"strings"
	"time"

	// protovalidate "buf.build/protovalidate-go"
	protovalidate "buf.build/go/protovalidate" //menggunakan ini
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func loggingMiddleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	log.Println("Masuk logging Middleware")
	log.Println(info.FullMethod)  //?melihat informasi method yang dipanggil dan usernya
	res, err := handler(ctx, req) //?jalan ke handler

	log.Println("Setelah request")
	return res, err
}

func authMiddleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	log.Println("Masuk auth Middleware")
	md, ok := metadata.FromIncomingContext(ctx) //?mengambil metadata (seperti token)
	if !ok {
		return nil, status.Error(codes.Unknown, "failed parsing metadata")
	}

	authToken, ok := md["authorization"]
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "token doesn't exist")
	}

	// log.Println(authToken[0])

	splitToken := strings.Split(authToken[0], " ") //?split tokennya
	token := splitToken[1]                         //?ambil tokennya

	if token != "secret" { //?jika tokennya bukan secret, (bisa diganti sesuai kebutuhan seperti jwt / session token)
		return nil, status.Error(codes.Unauthenticated, "token is not valid")
	}
	return handler(ctx, req) //?jalan ke handlerhandler(ctx, req) //?jalan ke handler
}

type userService struct {
	user.UnimplementedUserServiceServer //mengiinitialkan semua API USER (mungkin seperti resource di route laravel)
}

func (us *userService) CreateUser(ctx context.Context, userRequest *user.User) (*user.CreateResponse, error) {

	if err := protovalidate.Validate(userRequest); err != nil { //!akan masuk validasi ini ketika error
		if ve, ok := err.(*protovalidate.ValidationError); ok {
			var validations []*common.ValidationError = make([]*common.ValidationError, 0)
			for _, fieldErr := range ve.Violations {
				log.Printf("Field %s message %s", *fieldErr.Proto.Field.Elements[0].FieldName, *fieldErr.Proto.Message)

				validations = append(validations, &common.ValidationError{
					Field:   *fieldErr.Proto.Field.Elements[0].FieldName,
					Message: *fieldErr.Proto.Message,
				})
			}
			return &user.CreateResponse{
				Base: &common.BaseResponse{
					ValidationErrors: validations,
					StatusCode:       400,
					IsSuccess:        false,
					Message:          "There is validation error",
				},
			}, nil
		}
		return nil, status.Errorf(codes.InvalidArgument, "validation error %v", err)
	}
	// return nil, status.Errorf(codes.Internal, "Server is bugged") //?membuat example error internal

	log.Println("CreateUser is running")
	return &user.CreateResponse{
		Base: &common.BaseResponse{
			StatusCode: 0,
			IsSuccess:  true,
			Message:    "Success Create User",
		},
		CreatedAt: timestamppb.Now(),
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
	lis, err := net.Listen("tcp", ":8082") //harus sama dengan port di grpcclient/main.go
	if err != nil {
		log.Fatal("There is error in your net listen ", err)
	}

	//! setiap kali request, maka middleware ini akan dijalankan
	// serv := grpc.NewServer(grpc.ChainUnaryInterceptor(loggingMiddleware))
	serv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			loggingMiddleware, authMiddleware),
	)

	user.RegisterUserServiceServer(serv, &userService{})

	chat.RegisterChatServiceServer(serv, &chatService{})

	reflection.Register(serv)

	if err := serv.Serve(lis); err != nil {
		// log.Fatav...: "Error running server ", err)
		log.Fatal("Error running server ", err)
	}
}
