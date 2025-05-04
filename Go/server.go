package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"gRPC-Basic-Communication/pb"
)

type UserService struct {
	pb.UserServiceServer
}

func (s *UserService) GetUser(ctx context.Context, request *pb.DataRequest) (*pb.DataResponse, error) {
    age := int32(25)
    email := "helmi@xeno.com.my"

    return &pb.DataResponse{
        Id:    request.Id,
        Name:  "Helmi Aziz",
        Age:   &age,
        Email: &email,
    }, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50055")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer( grpcServer, &UserService{})
	fmt.Println( "gRPC Server running on port 50055" )

	if err := grpcServer.Serve( listener ); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	
}