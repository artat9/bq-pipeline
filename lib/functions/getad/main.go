package main

import (
	"aurora-backend/lib/functions/lib"
	"aurora-backend/lib/functions/lib/ad"
	"aurora-backend/lib/functions/lib/infrastructure/arweave"
	"aurora-backend/lib/functions/lib/infrastructure/eth/admanager"
	"context"
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ethereum/go-ethereum/common"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req, err := toInput(request)
	if err != nil {
		return unexpectedErrorAsNotFound(request, err), err
	}
	app, err := newApp(ctx)
	if err != nil {
		return unexpectedErrorAsNotFound(request, err), err
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
	metadata := request.PathParameters["metadata"]
	if account == "" {
		return req, errors.New("account required")
	}
	if metadata == "" {
		return req, errors.New("index required")
	}
	if !common.IsHexAddress(account) {
		return req, errors.New("account should be hex address")
	}
	return ad.GetInput{
		Account:  account,
		Metadata: metadata,
	}, nil
}

func newApp(ctx context.Context) (*ad.Service, error) {
	cli, err := arweave.NewClient()
	if err != nil {
		return &ad.Service{}, err
	}
	contract, err := admanager.NewProvider()
	if err != nil {
		return &ad.Service{}, err
	}
	return ad.NewWithContract(cli, contract), nil
}

func unexpectedErrorAsNotFound(request events.APIGatewayProxyRequest, err error) events.APIGatewayProxyResponse {
	return anyError(request, 404, err)
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
