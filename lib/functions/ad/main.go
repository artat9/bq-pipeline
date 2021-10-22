package main

import (
	"context"
	"encoding/json"
	"kaleido-backend/lib/functions/lib"
	"kaleido-backend/lib/functions/lib/ad"
	"kaleido-backend/lib/functions/lib/infrastructure/arweave"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req, err := toInput(request)
	if err != nil {
		return unexpectedError(request), err
	}
	app, err := newApp(ctx)
	if err != nil {
		return unexpectedError(request), err
	}
	res, err := app.Upload(ctx, req)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    lib.Headers(request),
			Body:       "unknown error",
		}, err
	}
	resJSON, _ := json.Marshal(&res)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    lib.Headers(request),
		Body:       string(resJSON),
	}, nil
}

func main() {
	lambda.Start(handler)
}

func toInput(request events.APIGatewayProxyRequest) (ad.Input, error) {
	req := ad.Input{}
	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		return ad.Input{}, err
	}
	return req, nil
}

func newApp(ctx context.Context) (*ad.Service, error) {
	cli, err := arweave.NewClient()
	if err != nil {
		return &ad.Service{}, err
	}
	return ad.NewService(cli), nil
}

func unexpectedError(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Headers:    lib.Headers(request),
		Body:       "unknown error",
	}
}
