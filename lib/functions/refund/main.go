package main

import (
	"aurora-backend/lib/functions/lib"
	"aurora-backend/lib/functions/lib/infrastructure/arweave"
	"aurora-backend/lib/functions/lib/refund"
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response response
type Response struct {
	Metadata string `json:"metadata"`
	Image    string `json:"image"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	app, err := newApp(ctx)
	if err != nil {
		return unexpectedError(), err
	}
	res, img, err := app.Upload(ctx)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    lib.Headers(),
			Body:       "unknown error",
		}, err
	}
	resJSON := Response{Metadata: res, Image: img}
	resp, err := json.Marshal(&resJSON)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    lib.Headers(),
			Body:       "unknown error",
		}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    lib.Headers(),
		Body:       string(resp),
	}, nil
}

func main() {
	lambda.Start(handler)
}

func newApp(ctx context.Context) (*refund.Service, error) {
	cli, err := arweave.NewClient()
	if err != nil {
		return &refund.Service{}, err
	}
	return refund.NewService(cli), nil
}

func unexpectedError() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Headers:    lib.Headers(),
		Body:       "unknown error",
	}
}
