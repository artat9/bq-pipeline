package main

import (
	"aurora-backend/lib/functions/lib"
	"aurora-backend/lib/functions/lib/ad"
	"aurora-backend/lib/functions/lib/infrastructure/arweave"
	"context"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req, err := toInput(request)
	if err != nil {
		return unexpectedError(request, err), err
	}
	app, err := newApp(ctx)
	if err != nil {
		return unexpectedError(request, err), err
	}
	res, err := app.Get(ctx, req)
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

func toInput(request events.APIGatewayProxyRequest) (ad.GetInput, error) {
	req := ad.GetInput{}
	account := request.PathParameters["account"]
	idx := request.PathParameters["index"]
	if account == "" {
		return req, errors.New("account required")
	}
	if idx == "" {
		return req, errors.New("index required")
	}
	i, err := strconv.Atoi(idx)
	if err != nil {
		return req, err
	}
	return ad.GetInput{
		Account: account,
		Index:   i,
	}, nil
}

func newApp(ctx context.Context) (*ad.Service, error) {
	cli, err := arweave.NewClient()
	if err != nil {
		return &ad.Service{}, err
	}
	return ad.NewService(cli), nil
}

func unexpectedError(request events.APIGatewayProxyRequest, err error) events.APIGatewayProxyResponse {
	return anyError(request, 500, err)
}

func anyError(request events.APIGatewayProxyRequest, code int, err error) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Headers:    lib.Headers(request),
		Body:       err.Error(),
	}
}

func clientError(request events.APIGatewayProxyRequest, err error) events.APIGatewayProxyResponse {
	return anyError(request, 400, err)
}
