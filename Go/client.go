package main

import (
	"context"
	"fmt"
	"log"
	// "github.com/davecgh/go-spew/spew"

	"google.golang.org/grpc"
	"gRPC-Basic-Communication/pb"
)

func main() {
	// Connect to gRPC server
		insecureOption := grpc.WithInsecure() // No SSL/TLS - Only for local development
		connection, err := grpc.Dial("localhost:50055", insecureOption )
		if err != nil {
			log.Fatal(err)
		}
		defer connection.Close()

	// Send request
		request := &pb.DataRequest{Id: "30"}
		client := pb.NewUserServiceClient( connection )
		response, err := client.GetUser(context.Background(), request)
		if err != nil {
			log.Fatal(err)
		}

	// Display response
		fmt.Printf("Receive response => \nID:%s\nName:%s\nEmail:%s\nAge:%d",
							response.Id,
							response.Name,
							*response.Email,
							response.GetAge() )
}