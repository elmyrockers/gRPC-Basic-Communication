package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"github.com/davecgh/go-spew/spew"

	"google.golang.org/grpc"
	"gRPC-Basic-Communication/pb"
)

type UserService struct {
	pb.UserServiceServer
}

func (s *UserService) GetUser(ctx context.Context, request *pb.DataRequest) (*pb.DataResponse, error) {
	// Debugging Request
		spew.Dump( request )


	// Response
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
	// Register the service
		grpcServer := grpc.NewServer()
		pb.RegisterUserServiceServer( grpcServer, &UserService{})

	// Start gRPC server
		// Listen on TCP port 50055
			listener, err := net.Listen("tcp", ":50055")
			if err != nil {
				panic(err)
			}

		// Start serving requests
			fmt.Println( "Starting gRPC server on port 50055..." )
			if err := grpcServer.Serve( listener ); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}
}