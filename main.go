package main

import (
	"grpc-course-protobuf/pb/user"
	"log"
	"net"

	"google.golang.org/grpc"
)

type userService struct {
	user.UnimplementedUserServiceServer //mengiinitialkan semua API USER (mungkin seperti resource di route laravel)
}

func main() {
	// lis, err := net.Listen(network: "tcp", address: ":8080") //ditutorial seperti ini
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("There is error in your net listen ", err)
	}

	serv := grpc.NewServer()

	user.RegisterUserServiceServer(serv, &userService{})

	if err := serv.Serve(lis); err != nil {
		// log.Fatal(v...: "Error running server ", err)
		log.Fatal("Error running server ", err)
	}
}
