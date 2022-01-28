package main

import (
	"bq-pipeline/pkg/handle"
	"bq-pipeline/pkg/infrastructure/ddb"
	"bq-pipeline/pkg/infrastructure/sns"
	"bq-pipeline/pkg/mediaaccount"
	mediarep "bq-pipeline/pkg/mediaaccount/persistence"
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event interface{}) (mediaaccount.ApplyForMediaOutput, error) {
	eoa, err := handle.EOAFromSign(event)
	if err != nil {
		return mediaaccount.ApplyForMediaOutput{}, err
	}
	res, err := mediaaccount.NewService(mediarep.New(ddb.New()), sns.New()).NewApplication(ctx, eoa, toInput(event))
	if err != nil {
		return mediaaccount.ApplyForMediaOutput{}, err
	}
	return res, nil
}

func toInput(event interface{}) mediaaccount.ApplyForMediaInput {
	return mediaaccount.ApplyForMediaInput{
		Name:             handle.Argument(event, "name"),
		MailAddress:      handle.Argument(event, "mailAddress"),
		URL:              handle.Argument(event, "url"),
		Description:      handle.Argument(event, "description"),
		PrimaryCustomers: handle.Argument(event, "primaryCustomers"),
		DocumentURL:      handle.Argument(event, "documentUrl"),
		PVMonth:          handle.IntArgument(event, "pvMonth"),
	}
}

func main() {
	lambda.Start(handler)
}
