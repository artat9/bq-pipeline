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

func (s SNS) NotifyApplicationCreated(ctx context.Context, application mediaaccount.Application) error {
	input := fromApplication(application)
	payload, err := json.Marshal(&input)
	if err != nil {
		log.Error("json unmarshal failed", err)
		return err
	}
	s.svc.PublishWithContext(ctx, &sns.PublishInput{
		Message: aws.String(string(payload)),
		Subject: aws.String(string(application.Account.Hex())),
	})
	return nil
}

func FromEvent(input events.SNSEvent) ([]mediaaccount.Application, error) {
	res := []mediaaccount.Application{}
	for _, r := range input.Records {
		in := mediaaccount.Application{}
		payload, err := json.Marshal(&r)
		if err != nil {
			log.Error("json marshal failed", err)
			return res, err
		}
		if err = json.Unmarshal(payload, &in); err != nil {
			log.Error("json unmarshal failed", err)
			return res, err
		}
		res = append(res, in)
	}
	return res, nil
}

func newMediaAppliedTopicArn() string {
	topicName := os.Getenv("ApplicationCreatedTopicName")
	return "arn:aws:sns:us-east-1:495476032358:" + topicName
}
