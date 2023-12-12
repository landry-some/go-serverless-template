package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/gkatanacio/go-serverless-template/pkg/greeting"
)

var greetingService greeting.Servicer

func init() {
	greetingService = greeting.NewService()
}

func handler(ctx context.Context, request *events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	log.Printf("%+v", request)

	msg := greetingService.HelloMessage(os.Getenv("GREET_WHO"))

	return &events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: fmt.Sprintf(`{"message":"%s"}`, msg),
	}, nil
}

func main() {
	lambda.Start(handler)
}
