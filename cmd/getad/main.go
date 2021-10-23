package main

import (
	"context"
	"errors"
	"kaleido-backend/pkg/ad"
	"kaleido-backend/pkg/common/log"
	"kaleido-backend/pkg/handle"
	"kaleido-backend/pkg/infrastructure/arweave"
	"kaleido-backend/pkg/infrastructure/eth/admanager"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ethereum/go-ethereum/common"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req, err := toInput(request)
	if err != nil {
		log.Error("unexpected error", err)
		return handle.NotFoundErrorResponse(request, err), nil
	}
	app, err := newApp(ctx)
	if err != nil {
		log.Error("unexpected err", err)
		return handle.NotFoundErrorResponse(request, err), err
	}
	res, err := app.Get(ctx, req)
	if err != nil {
		return handle.UnexpectedErrorResponse(request), err
	}
	return handle.NormalResponse(request, res), err
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
		return req, errors.New("metadata required")
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
