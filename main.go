package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is our Lambda function handler
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	html := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>My Go Lambda Page</title>
    </head>
    <body>
        <h1>Welcome to My Go Lambda-Rendered Page!</h1>
        <p>This page is served by an AWS Lambda function written in Go.</p>
    </body>
    </html>
    `

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
