package main

import (
	"context"
	"fmt"
	"time"

	"github.com/journeyai/grpc-validation/server/protocols/example"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := example.NewExampleServiceClient(conn)

	_, err = client.CreateWidget(ctx, &example.CreateWidgetRequest{
		Name:        "",
		Age:         18,
		Email:       "joe",
		PhoneNumber: "(303) 555-1234",
		Attributes:  []string{"asd", ""},
	})

	status, _ := status.FromError(err)
	fmt.Printf("\nCode: %s\n", status.Code().String())
	fmt.Printf("Message: %s\n\n", status.Message())

	details := status.Details()
	badRequest := details[0].(*errdetails.BadRequest)
	for _, violation := range badRequest.FieldViolations {
		fmt.Printf("  %s: %s\n", violation.Field, violation.Description)
	}
	fmt.Println("")

}
