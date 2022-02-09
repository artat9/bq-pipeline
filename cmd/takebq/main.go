package main

import (
	"bq-pipeline/pkg/common/log"
	"bq-pipeline/pkg/infrastructure/bigquery"
	"bq-pipeline/pkg/infrastructure/ssm"
	"fmt"

	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.SNSEvent) error {
	client, err := bigquery.NewClient(ctx, ssm.New())

	if err != nil {
		log.Error("bigquery client create failed", err)
		return err
	}
	defer client.Close()

	reader := bigquery.NewReader(ctx, client)
	service := bigquery.NewService(nil, reader)
	users, err := service.Download(ctx)

	if err != nil {
		log.Error("fetch users data failed", err)
		return err
	}
	fmt.Printf("users: %+v\n", users)
	return nil
}

func main() {
	lambda.Start(handler)
}
