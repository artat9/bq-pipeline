package main

import (
	"context"
	"kaleido-backend/pkg/handle"
	"kaleido-backend/pkg/image"
	"kaleido-backend/pkg/infrastructure/s3"

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
