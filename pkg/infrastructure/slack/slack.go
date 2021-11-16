package slack

import (
	"context"
	"fmt"
	"kaleido-backend/pkg/common/log"
	"kaleido-backend/pkg/email"
	"kaleido-backend/pkg/mediaaccount"
	"os"
	"reflect"

	"github.com/slack-go/slack"
)

type (
	// Slack slack
	Slack struct {
		svc *slack.Client
	}
	// TokenResolver slack token resolver
	TokenResolver interface {
		SlackToken(ctx context.Context) (string, error)
	}
)

// New new service
func New(ctx context.Context, r TokenResolver) (Slack, error) {
	tk, err := r.SlackToken(ctx)
	if err != nil {
		return Slack{}, err
	}
	return Slack{
		svc: slack.New(tk),
	}, nil
}

func (s Slack) NotifyEmailAddressConfirmationCompleted(ctx context.Context, in email.VerificationInput) error {
	return s.notify(ctx, in.SignatureInput)
}

func (s Slack) notify(ctx context.Context, input interface{}) error {
	vs := reflect.ValueOf(input)
	txtblocks := []*slack.TextBlockObject{}
	for i := 0; i < vs.NumField(); i++ {
		txtblocks = append(txtblocks, slack.NewTextBlockObject("plain_text", fmt.Sprintf("%v", vs.Field(i).Interface()), false, false))
	}
	_, _, err := s.svc.PostMessage(channelName(), slack.MsgOptionBlocks(
		slack.NewSectionBlock(
			&slack.TextBlockObject{Type: "mrkdwn", Text: "New Application arrival"},
			txtblocks,
			nil,
		),
	))
	if err != nil {
		log.Error("unexpected error", err)
	}
	return err
}

// NotifyApplicationCreated notify to a slack channel
func (s Slack) NotifyApplicationCreated(ctx context.Context, application mediaaccount.Application) error {
	return s.notify(ctx, application)
}

func channelName() string {
	prefix := "#kaleido-media-申請-"
	env := os.Getenv("EnvironmentId")
	if env == "v1dev" {
		env = "dev"
	}
	return prefix + env
}
