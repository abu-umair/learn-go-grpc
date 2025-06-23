package main

import (
	"context"
	"grpc-course-protobuf/pb/user"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type userService struct {
	user.UnimplementedUserServiceServer //mengiinitialkan semua API USER (mungkin seperti resource di route laravel)
}

func (us *userService) CreateUser(ctx context.Context, userRequest *user.User) (*user.CreateResponse, error) {
	log.Println("CreateUser is running")
	return &user.CreateResponse{
		Message: "Success Create User",
	}, nil
}

func main() {
	// lis, err := net.Listen(network: "tcp", address: ":8080") //ditutorial seperti ini
	lis, err := net.Listen("tcp", ":8081") //harus sama dengan port di grpcclient/main.go
	if err != nil {
		log.Fatal("There is error in your net listen ", err)
	}

	serv := grpc.NewServer()

	user.RegisterUserServiceServer(serv, &userService{})

	reflection.Register(serv)

	if err := serv.Serve(lis); err != nil {
		// log.Fatal(v...: "Error running server ", err)
		log.Fatal("Error running server ", err)
	}
}
