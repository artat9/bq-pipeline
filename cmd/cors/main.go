package main

import (
	"bq-pipeline/pkg/handle"
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return handle.NormalResponse(request, ""), nil
}

func main() {
	lambda.Start(handler)
}
