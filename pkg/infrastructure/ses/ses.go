package ses

import (
	"bq-pipeline/pkg/common/log"
	"bq-pipeline/pkg/email"
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type (
	// SES ses service
	SES struct {
		svc service
	}

	// service service
	service interface {
		SendEmailWithContext(aws.Context, *ses.SendEmailInput, ...request.Option) (*ses.SendEmailOutput, error)
	}
)

// New new service
func New() SES {
	return SES{
		svc: ses.New(session.New(), &aws.Config{Region: aws.String("us-east-1")}),
	}
}

// SendApplicationCreated send application created email
func (s SES) SendApplicationCreated(ctx context.Context, in email.VerificationInput) error {
	_, err := s.svc.SendEmailWithContext(ctx, &ses.SendEmailInput{
		Destination: &ses.Destination{ToAddresses: aws.StringSlice([]string{in.MailAddress})},
		Message: &ses.Message{
			Subject: content("[Kaleido] Email Verification"),
			Body: &ses.Body{
				Text: content("To complete media application, click the following link: " + url(in)),
			},
		},
		Source: aws.String(mailFrom()),
	})
	if err != nil {
		log.Error("ses:sendEmail error", err)
	}
	return err
}

func content(text string) *ses.Content {
	return &ses.Content{
		Charset: aws.String("UTF-8"),
		Data:    aws.String(text),
	}
}

func mailFrom() string {
	rootDomain := os.Getenv("RootDomain")
	return "no-reply@" + rootDomain
}

func url(in email.VerificationInput) string {
	return fmt.Sprintf("https://%s/callback/verification/email?%s", os.Getenv("RootDomain"), in.QueryString())
}
