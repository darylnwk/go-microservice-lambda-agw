package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-xray-sdk-go/xray"
)

func init() {
	// Initialise variables or configurations here to be used across all Lambda instances
}

func main() {
	lambda.Start(func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		// Initialise AWS X-Ray
		xray.Configure(xray.Config{})

		return events.APIGatewayProxyResponse{}, nil
	})
}
