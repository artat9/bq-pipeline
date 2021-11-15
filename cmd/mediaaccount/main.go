package main

import (
	"context"
	"kaleido-backend/pkg/handle"
	"kaleido-backend/pkg/infrastructure/ddb"
	"kaleido-backend/pkg/infrastructure/slack"
	"kaleido-backend/pkg/infrastructure/ssm"
	"kaleido-backend/pkg/mediaaccount"
	mediarep "kaleido-backend/pkg/mediaaccount/persistence"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event interface{}) (*mediaaccount.ApplyForMediaOutput, error) {
	eoa, err := handle.EOAFromSign(event)
	if err != nil {
		return nil, err
	}
	slc, err := slack.New(ctx, ssm.New())
	res, err := mediaaccount.NewService(mediarep.New(ddb.New()), slc).NewApplication(ctx, eoa, toInput(event))
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func toInput(event interface{}) mediaaccount.ApplyForMediaInput {
	name := handle.Argument(event, "name")
	mail := handle.Argument(event, "mailAddress")
	url := handle.Argument(event, "url")
	description := handle.Argument(event, "description")
	return mediaaccount.ApplyForMediaInput{
		Name:        name,
		MailAddress: mail,
		URL:         url,
		Description: description,
	}
}

func main() {
	lambda.Start(handler)
}
