package main

import (
	"bq-pipeline/pkg/ad"
	"bq-pipeline/pkg/common/log"
	"bq-pipeline/pkg/handle"
	"bq-pipeline/pkg/infrastructure/eth/admanager"
	"bq-pipeline/pkg/infrastructure/ssm"
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request interface{}) (ad.GetOutput, error) {
	req, err := toInput(request)
	if err != nil {
		log.Error("unexpected error", err)
		return ad.GetOutput{}, err
	}
	app, err := newApp(ctx)
	if err != nil {
		log.Error("unexpected err", err)
		return ad.GetOutput{}, err
	}
	return app.Get(ctx, req)
}

func main() {
	lambda.Start(handler)
}

func toInput(request interface{}) (ad.GetInput, error) {
	spaceMetadata := handle.Argument(request, "spaceMetadata")
	return ad.GetInput{
		SpaceMetadata: spaceMetadata,
	}, nil
}

func newApp(ctx context.Context) (*ad.Service, error) {
	contract, err := admanager.NewProvider(ctx, ssm.New())
	if err != nil {
		return &ad.Service{}, err
	}
	return ad.NewWithContract(contract), nil
}
