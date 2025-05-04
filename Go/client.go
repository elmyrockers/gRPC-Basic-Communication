package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"gRPC-Basic-Communication/pb"
)

func main() {
	insecureOption := grpc.WithInsecure()
	connection, err := grpc.Dial("localhost:50055", insecureOption )
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient( connection )
	request := &pb.DataRequest{Id: "30"}

	response, err := client.GetUser(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive response => %s ", response.Name)
}