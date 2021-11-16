package main

import (
	"context"
	"kaleido-backend/pkg/common/log"
	"kaleido-backend/pkg/infrastructure/slack"
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
	sl, err := slack.New(ctx, ssm.New())
	if err != nil {
		return err
	}
	for _, app := range applications {
		if err = sl.NotifyApplicationCreated(ctx, app); err != nil {
			log.Error("slack notification failed", err)
			return err
		}
	}
	return nil
}
func main() { lambda.Start(handler) }

//func main() {
//	os.Setenv("EnvironmentId", "v1dev")
//	sl, err := slack.New(context.Background(), ssm.New())
//	if err != nil {
//		os.Exit(1)
//	}
//	if err = sl.NotifyApplicationCreated(context.Background(), mediaaccount.Application{
//		MailAddress: "test@example.com",
//		Account:     common.HexToAddress(""),
//		Name:        "sample",
//		URL:         "https://test.example.com",
//		Description: "this is test",
//	}); err != nil {
//		log.Error("slack notification failed", err)
//		os.Exit(1)
//	}
//}
//
