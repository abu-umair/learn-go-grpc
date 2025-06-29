package main

import (
	"context"
	"grpc-course-protobuf/pb/user"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	clientConn, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials())) //insecure.NewCredentials(): jika dilingkungan develop
	if err != nil {
		log.Fatal("Failed create client", err)
	}

	userClient := user.NewUserServiceClient(clientConn)

	res, err := userClient.CreateUser(context.Background(), &user.User{
		// Age: -1, //?inputnya negatif, jadi error
		Age: 20, //?inputnya positif, jadi tidak ada error
	})
	if err != nil {
		st, ok := status.FromError(err)
		// error gRPC
		if ok { //? ok akan true jika errornya dari gRPC ( false jika error selain gRPC )
			if st.Code() == codes.InvalidArgument { //?apakah errornya  == invalid argument
				log.Println("There is validation error: ", st.Message()) //?message: bisa kita gunakan untuk menampilkan pesan error
			} else if st.Code() == codes.Unknown {
				log.Println("There is unknown error: ", st.Message())
			} else if st.Code() == codes.Internal {
				log.Println("There is internal error: ", st.Message())
			}
			return
		}

		log.Println("Failed to send message", err) //?menganggap responsenya error
		log.Println(res)                           //? nilainya nil / gak ada response
		return
	}

	log.Println("Response from server ", res.Message) //?response dari server

}
