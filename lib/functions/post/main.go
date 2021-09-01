package main

import (
	"aurora-backend/lib/functions/lib"
	"aurora-backend/lib/functions/lib/infrastructure/arweave"
	"aurora-backend/lib/functions/lib/post"
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req, err := toInput(request)
	if err != nil {
		return unexpectedError(), err
	}
	app, err := newApp(ctx)
	if err != nil {
		return unexpectedError(), err
	}
	res, err := app.Upload(ctx, req)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    lib.Headers(),
			Body:       "unknown error",
		}, err
	}
	resJSON, _ := json.Marshal(&res)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    lib.Headers(),
		Body:       string(resJSON),
	}, nil
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

func unexpectedError() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Headers:    lib.Headers(),
		Body:       "unknown error",
	}
}