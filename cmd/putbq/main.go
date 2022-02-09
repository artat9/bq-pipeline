package main

import (
	"bq-pipeline/pkg/common/log"
	"bq-pipeline/pkg/infrastructure/bigquery"
	"bq-pipeline/pkg/infrastructure/ssm"

	"bq-pipeline/pkg/user"
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.SNSEvent) error {
	//TODO FIX This is sample data.
	sampleUser := user.User{ID: "1", Name: "hoge"}
	users := []user.User{}
	users = append(users, sampleUser)

	client, err := bigquery.NewClient(ctx, ssm.New())
	if err != nil {
		log.Error("bigquery client create failed", err)
		return err
	}
	defer client.Close()

	writer := bigquery.NewWriter(client)
	bq := bigquery.NewService(writer, nil)
	svc := user.NewService(bq, nil)

	if err = svc.Upload(ctx, users); err != nil {
		log.Error("upload failed", err)
		return err
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
