package main

import (
	"bq-pipeline/pkg/common/log"
	"bq-pipeline/pkg/infrastructure/bigquery"
	"bq-pipeline/pkg/infrastructure/ssm"
	"fmt"
	"os"

	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.SNSEvent) error {
	client, err := bigquery.NewClient(ctx, ssm.New())

	if err != nil {
		log.Error("bigquery client create failed", err)
		os.Exit(1)
	}
	defer client.Close()

	sr := bigquery.NewSampleReader(ctx, client)
	service := bigquery.NewService(nil, sr)
	users, err := service.Download(ctx)

	if err != nil {
		log.Error("fetch users data failed", err)
		return err
	}
	fmt.Println(users)
	return nil
}

func main() {
	lambda.Start(handler)
}
