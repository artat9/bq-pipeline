package main

import (
	"context"
	"kaleido-backend/pkg/infrastructure/cloudfront"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.S3Event) error {
	tar := []string{}
	for _, e := range event.Records {
		tar = append(tar, e.S3.Object.Key)
	}
	return cloudfront.New().Invalidate(ctx, tar)
}

func main() {
	lambda.Start(handler)
}
