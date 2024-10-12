package main

import (
	"context"

	"subway-lambda/internal/common"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "text/plain"},
		Body:            common.RandomStation(),
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
