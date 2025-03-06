package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/gkatanacio/go-serverless-template/pkg/greeting"
)

var greetingService greeting.Service

func init() {
	greetingService = greeting.NewService()
}

func handle(ctx context.Context, request *events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	slog.InfoContext(ctx, fmt.Sprintf("Received request: %+v", request))

	msg := greetingService.HelloMessage(ctx, os.Getenv("GREET_WHO"))

	return &events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: fmt.Sprintf(`{"message":"%s"}`, msg),
	}, nil
}

func main() {
	lambda.Start(handle)
}
