package main

import (
	"bytes"
	"context"

	"subway-lambda/internal/common"
	"subway-lambda/internal/templ"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var buf bytes.Buffer
	component := templ.Home(common.RandomStation())
	component.Render(context.Background(), &buf)
	html := buf.String()

	return events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "text/html"},
		Body:            html,
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
