package main

import (
	"aurora-backend/lib/functions/lib"
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    lib.Headers(request),
	}, nil
}

func main() {
	lambda.Start(handler)
}
