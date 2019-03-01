package main

import (
	"fmt"
	"log"
	"net"

	"github.com/code-sleuth/grpc-go/greet/greetpb"
	"google.golang.org/grpc"
)

// server struct
type server struct{}

func main() {
	fmt.Println("Hello Ibra")

	// create listner lis
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	//create grpc server
	s := grpc.NewServer()

	//register service(s) to grpc server
	greetpb.RegisterGreetServiceServer(s, &server{})

	// bind port to grpc server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
