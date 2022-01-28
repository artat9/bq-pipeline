package main

import (
	"bq-pipeline/pkg/handle"
	"bq-pipeline/pkg/image"
	"bq-pipeline/pkg/infrastructure/s3"
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event interface{}) (*image.UpdateOutput, error) {
	filename := handle.Argument(event, "filename")
	eoa, err := handle.EOAFromSign(event)
	if err != nil {
		return nil, err
	}
	service := image.New(s3.New())
	res, err := service.Update(ctx, eoa, image.UpdateInput{Name: filename})
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func main() {
	lambda.Start(handler)
}
