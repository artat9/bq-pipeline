package main

import (
	"context"
	"encoding/json"
	"kaleido-backend/pkg/handle"
	"kaleido-backend/pkg/infrastructure/arweave"
	"kaleido-backend/pkg/post"

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
		handle.UnexpectedErrorResponse(request)
	}
	return handle.NormalResponse(request, res), err
}

func main() {
	lambda.Start(handler)
}

func toInput(request events.APIGatewayProxyRequest) (post.Input, error) {
	req := post.Input{}
	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		return post.Input{}, err
	}
	return req, nil
}

func newApp(ctx context.Context) (*post.Service, error) {
	cli, err := arweave.NewClient()
	if err != nil {
		return &post.Service{}, err
	}
	return post.NewService(cli), nil
}
