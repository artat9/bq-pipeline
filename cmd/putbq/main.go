package main

import (
	"bq-pipeline/pkg/email"
	"bq-pipeline/pkg/infrastructure/ses"
	signer "bq-pipeline/pkg/infrastructure/sign"
	"bq-pipeline/pkg/infrastructure/sns"
	"bq-pipeline/pkg/infrastructure/ssm"
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.SNSEvent) error {
	applications, err := sns.FromEvent(request)
	if err != nil {
		return err
	}
	signer, err := signer.NewWithResolver(context.Background(), ssm.New())
	if err != nil {
		return err
	}
	for _, app := range applications {
		if err = email.NewSignService(signer, ses.New()).SendEmailVerification(ctx, app.MailAddress); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
