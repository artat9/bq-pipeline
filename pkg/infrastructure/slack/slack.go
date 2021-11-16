package slack

import (
	"context"
	"kaleido-backend/pkg/common/log"
	"kaleido-backend/pkg/mediaaccount"
	"os"

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

// NotifyApplicationCreated notify to a slack channel
func (s Slack) NotifyApplicationCreated(ctx context.Context, application mediaaccount.Application) error {
	_, _, err := s.svc.PostMessage(channelName(), slack.MsgOptionBlocks(
		slack.NewSectionBlock(
			&slack.TextBlockObject{Type: "mrkdwn", Text: "New Application arrival"},
			[]*slack.TextBlockObject{
				{Type: "plain_text", Text: application.Name},
				{Type: "plain_text", Text: application.URL},
				{Type: "plain_text", Text: application.Description},
				{Type: "plain_text", Text: application.Account.Hex()},
				{Type: "plain_text", Text: application.MailAddress},
			},
			nil,
		),
	))
	if err != nil {
		log.Error("unexpected error", err)
	}
	return err
}

func channelName() string {
	prefix := "#kaleido-media-申請-"
	env := os.Getenv("EnvironmentId")
	if env == "v1dev" {
		env = "dev"
	}
	return prefix + env
}
