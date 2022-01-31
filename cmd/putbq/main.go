package main

import (
	"bq-pipeline/pkg/common/log"
	"bq-pipeline/pkg/infrastructure/bigquery"
	"bq-pipeline/pkg/infrastructure/ssm"
	"os"

	"bq-pipeline/pkg/user"
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.SNSEvent) error {
	//TODO FIX 仮のデータとしてインプットデータを生成
	sampleUser := user.User{ID: "1", Name: "hoge"}
	users := []user.User{}
	users = append(users, sampleUser)

	client, err := bigquery.NewClient(ctx, ssm.New())
	if err != nil {
		log.Error("bigquery client create failed", err)
		os.Exit(1)
	}
	defer client.Close()

	si := bigquery.NewSampleInserter(ctx, users, client)
	bq := bigquery.NewService(si, nil)
	svc := user.NewService(bq, nil)

	if err = svc.Upload(ctx); err != nil {
		log.Error("upload failed", err)
		os.Exit(1)
	}

	return nil

	// i := user.Service{}

	//service := user.New(bigquery.NewSampleInserter(ctx, users))

	// user.Service{
	// 	bigquery.NewSampleInserter(ctx, inputUserData),
	// }
	// inserter := bigquery.NewSampleInserter(ctx, inputUserData)

	// hoge := bigquery.Service{bigquery.NewSampleInserter(ctx, inputUserData)}

	//service2 := user.Service{bigquery.New(ctx).Upload(ctx,)}
	//hoge := service2.Upload(ctx, inputUserData)
	//service := bigquery.Service{}

	//err := service2.Upload(ctx, inputUserData)
	// if err != nil {
	// 	return err
	// }

	// service2.Upload(mediarep.New(ddb.New()), sns.New()).NewApplication(ctx, eoa, toInput(event))
	//client, err := bigquery.New(ctx)

	// if err != nil {
	// 	log.Error("bigquery client create failed", err)
	// 	os.Exit(1)
	// }
	// defer client.Close()

	// service := bigquery.
	//bq := bigquery.Service{bigquery.NewSampleInserter()}

	//service := user.Service{bigquery.NewService()}

	// if err = user.Service{bigquery.Service{}}.Upload(ctx, client, inputUserData); err != nil {
	// 	return nil
	// }

	// i := &bigquery.Service{bigquery.NewSampleInserter(ctx, users, *client)}
	// b := bigquery.NewService()

	// if err := service.Upload(ctx, client, users); err != nil {
	// 	return err
	// }
	// return nil
}

func main() {
	lambda.Start(handler)
}
