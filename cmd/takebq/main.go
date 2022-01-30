package main

import (
	"bq-pipeline/pkg/common/log"
	"bq-pipeline/pkg/infrastructure/bigquery"
	"fmt"
	"os"

	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.SNSEvent) error {
  // インプットデータを生成
	service := bigquery.Service{}
	client, err := bigquery.New(ctx)

	if err != nil {
		log.Error("bigquery client create failed", err)
		os.Exit(1)
	}

	query := "SELECT " +
	         "id, name " +
					 "FROM " + os.Getenv("TARGET_GCP_PROJECT_ID") + "." + os.Getenv("TARGET_DATASET_ID") + ".sample_terraform_user_table"
	data, err := service.Download(ctx, *client, query)

  if err != nil {
		return err
	}
	fmt.Println(data)
	return nil
}

func main() {
	lambda.Start(handler)
}
