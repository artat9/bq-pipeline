package sns

import (
	"context"
	"encoding/json"
	"kaleido-backend/pkg/common/log"
	"kaleido-backend/pkg/mediaaccount"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/ethereum/go-ethereum/common"
)

type (
	// SNS sns
	SNS struct {
		svc snsiface.SNSAPI
	}
	// Payload payload
	Payload struct {
		mediaaccount.Application
		EOA string `json:"eoa"`
	}
)

// New new s3
func New() SNS {
	cli := sns.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))
	//xray.AWS(cli.Client)
	return SNS{
		svc: cli,
	}
}

func fromApplication(ac mediaaccount.Application) Payload {
	return Payload{
		Application: ac,
		EOA:         ac.Account.Hex(),
	}
}
func (p Payload) toApplication() mediaaccount.Application {
	p.Account = common.HexToAddress(p.EOA)
	return p.Application
}

// NotifyApplicationCreated notify application created
func (s SNS) NotifyApplicationCreated(ctx context.Context, application mediaaccount.Application) error {
	return s.notifyApplication(ctx, application, newMediaAppliedTopicArn())
}

func (s SNS) notifyApplication(ctx context.Context, application mediaaccount.Application, topicArn string) error {
	input := fromApplication(application)
	payload, err := json.Marshal(&input)
	if err != nil {
		log.Error("json unmarshal failed", err)
		return err
	}
	_, err = s.svc.PublishWithContext(ctx, &sns.PublishInput{
		Message:  aws.String(string(payload)),
		TopicArn: aws.String(topicArn),
	})
	if err != nil {
		log.Error("sns publish failed", err)
	}
	return err
}

// NotifyApplicationCompleted notify application completed
func (s SNS) NotifyApplicationCompleted(ctx context.Context, application mediaaccount.Application) error {
	return s.notifyApplication(ctx, application, completeMediaAppliedTopicArn())
}

// FromEvent from sns event to application
func FromEvent(input events.SNSEvent) ([]mediaaccount.Application, error) {
	res := []mediaaccount.Application{}
	for _, r := range input.Records {
		in := Payload{}
		if err := json.Unmarshal([]byte(r.SNS.Message), &in); err != nil {
			log.Error("json unmarshal failed", err)
			return res, err
		}
		in.Application.Account = common.HexToAddress(in.EOA)
		res = append(res, in.Application)
	}
	return res, nil
}

func newMediaAppliedTopicArn() string {
	return topicArn(os.Getenv("ApplicationCreatedTopicName"))
}

func completeMediaAppliedTopicArn() string {
	return topicArn(os.Getenv("ApplicationCompletedTopicName"))
}

func topicArn(suf string) string {
	return topicArnPrefix() + suf
}

func topicArnPrefix() string {
	return "arn:aws:sns:us-east-1:495476032358:"
}
