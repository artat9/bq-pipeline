package main

import (
	"context"
	"kaleido-backend/lib/functions/lib"

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
