package main

import (
	"context"
	"kaleido-backend/pkg/email"
	"kaleido-backend/pkg/infrastructure/ses"
	signer "kaleido-backend/pkg/infrastructure/sign"
	"kaleido-backend/pkg/infrastructure/sns"
	"kaleido-backend/pkg/infrastructure/ssm"

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
		if err = email.NewNotificationService(signer, ses.New()).SendEmailVerification(ctx, app.MailAddress); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
