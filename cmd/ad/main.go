package main

import (
	"context"
	"encoding/json"
	"kaleido-backend/pkg/ad"
	"kaleido-backend/pkg/handle"
	"kaleido-backend/pkg/infrastructure/arweave"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req, err := toInput(request)
	if err != nil {
		return handle.UnexpectedErrorResponse(request), err
	}
	app, err := newApp(ctx)
	if err != nil {
		return handle.UnexpectedErrorResponse(request), err
	}
	res, err := app.Upload(ctx, req)
	if err != nil {
		return handle.UnexpectedErrorResponse(request), err
	}
	return handle.NormalResponse(request, res), nil
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
