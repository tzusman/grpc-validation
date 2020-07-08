package main

import (
	"fmt"
	"net"

	"github.com/journeyai/grpc-validation/server/protocols/example"
	"google.golang.org/grpc"
)

func main() {

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	exampleServer, err := example.NewServer()
	if err != nil {
		panic(err)
	}
	example.RegisterExampleServiceServer(grpcServer, *exampleServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	fmt.Println("Service GRPC on 50051")
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
