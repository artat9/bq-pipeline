package main

import (
	"aurora-backend/lib/functions/lib"
	"aurora-backend/lib/functions/lib/common/log"
	"aurora-backend/lib/functions/lib/infrastructure/arweave"
	"aurora-backend/lib/functions/lib/infrastructure/eth/postmanager"
	"aurora-backend/lib/functions/lib/receipt"
	"aurora-backend/lib/functions/lib/receipt/generator"
	persreceipt "aurora-backend/lib/functions/lib/receipt/persistence"
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Metadata string `json:"metadata"`
	Image    string `json:"image"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req, err := toInput(request)
	if err != nil {
		return unexpectedError(), err
	}
	app, err := newApp(ctx)
	if err != nil {
		return unexpectedError(), err
	}
	res, img, err := app.Issue(ctx, req)
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

func toInput(request events.APIGatewayProxyRequest) (receipt.Input, error) {
	req := receipt.Input{}
	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		log.Error("marshal input failed", err)
		return receipt.Input{}, err
	}
	return req, nil
}

func newApp(ctx context.Context) (*receipt.Service, error) {
	cli, err := arweave.NewClient()
	if err != nil {
		return &receipt.Service{}, err
	}
	postmanager, err := postmanager.NewProvider()
	if err != nil {
		return &receipt.Service{}, err
	}
	receipt := receipt.NewService(persreceipt.New(cli), postmanager, generator.NewGenerator(), cli)
	return &receipt, nil
}

func unexpectedError() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Headers:    lib.Headers(),
		Body:       "unknown error",
	}
}
