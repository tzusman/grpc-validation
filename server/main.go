package main

import (
	"fmt"
	"net"

	p "github.com/journeyai/grpc-validation/protocols/example"
	"github.com/journeyai/grpc-validation/server/example"
	"google.golang.org/grpc"
)

func main() {

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	exampleServer, err := example.NewServer()
	if err != nil {
		panic(err)
	}
	p.RegisterExampleServiceServer(grpcServer, *exampleServer)

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
