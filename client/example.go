package main

import (
	"context"
	"fmt"
	"time"

	"github.com/journeyai/grpc-validation/protocols/example"
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

	/****************************************/
	/**   Poorly formatted Create Widget   **/
	/****************************************/

	_, err = client.CreateWidget(ctx, &example.CreateWidgetRequest{
		Name:        "",
		Age:         18,
		Email:       "joe",
		PhoneNumber: "(303) 555-1234",
		Attributes:  []string{"asd", ""},
	})

	s, _ := status.FromError(err)
	fmt.Printf("\nCode: %s\n", s.Code().String())
	fmt.Printf("Message: %s\n\n", s.Message())

	details := s.Details()
	badRequest := details[0].(*errdetails.BadRequest)
	for _, violation := range badRequest.FieldViolations {
		fmt.Printf("  %s: %s\n", violation.Field, violation.Description)
	}

	/****************************************/

	fmt.Println("\n\n~~~~~~~~~~~~~~~~~~~\n\n")

	/*************************************/
	/**   Poorly formatted ID request   **/
	/*************************************/

	_, err = client.GetWidget(ctx, &example.GetWidgetRequest{
		Id: "abcdef",
	})

	s, _ = status.FromError(err)
	fmt.Printf("\nCode: %s\n", s.Code().String())
	fmt.Printf("Message: %s\n\n", s.Message())

	details = s.Details()
	badRequest = details[0].(*errdetails.BadRequest)
	for _, violation := range badRequest.FieldViolations {
		fmt.Printf("  %s: %s\n", violation.Field, violation.Description)
	}

	/*************************************/

	fmt.Println("\n\n~~~~~~~~~~~~~~~~~~~\n\n")

	/************************************/
	/**   Create Widget Successfully   **/
	/************************************/

	createOut, err := client.CreateWidget(ctx, &example.CreateWidgetRequest{
		Name:        "Joe Shmoe",
		Age:         35,
		Email:       "joe@example.com",
		PhoneNumber: "+13035551234",
		Attributes:  []string{"friendly"},
	})

	fmt.Printf("Id:\t\t%s\n", createOut.Id)

	/************************************/

	fmt.Println("\n\n~~~~~~~~~~~~~~~~~~~\n\n")

	/********************/
	/**   Get Widget   **/
	/********************/

	getOut, err := client.GetWidget(ctx, &example.GetWidgetRequest{
		Id: createOut.Id,
	})

	fmt.Printf("Id:\t\t%s\n", getOut.Id)
	fmt.Printf("Name:\t\t%s\n", getOut.Name)
	fmt.Printf("Age:\t\t%d\n", getOut.Age)
	fmt.Printf("Email:\t\t%s\n", getOut.Email)
	fmt.Printf("PhoneNumber:\t%s\n", getOut.PhoneNumber)
	fmt.Printf("Attributes:\t%s\n", getOut.Attributes)

	/********************/

}
