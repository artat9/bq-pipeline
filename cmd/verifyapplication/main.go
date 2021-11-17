package main

import (
	"context"
	"kaleido-backend/pkg/common/log"
	"kaleido-backend/pkg/email"
	"kaleido-backend/pkg/handle"
	signer "kaleido-backend/pkg/infrastructure/sign"
	"kaleido-backend/pkg/infrastructure/slack"
	"kaleido-backend/pkg/infrastructure/ssm"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event interface{}) (Output, error) {
	verifier, err := signer.NewVerifierWithResolver(ctx, ssm.New())
	if err != nil {
		log.Error("initialization failed", err)
		return Output{}, err
	}
	sl, err := slack.New(ctx, ssm.New())
	if err != nil {
		log.Error("initialization failed", err)
		return Output{}, err
	}
	got, err := email.NewVerificationService(verifier, sl).VerifyEmail(ctx, toInput(event))
	if err != nil {
		log.Error("verification failed", err)
		return Output{}, err
	}
	return Output{MailAddress: got}, nil
}

// Output output
type Output struct {
	MailAddress string `json:"mailAddress"`
}

func toInput(event interface{}) email.VerificationInput {
	valstr := handle.Argument(event, "validity")
	val, _ := strconv.Atoi(valstr)
	return email.VerificationInput{
		SignatureInput: email.SignatureInput{
			MailAddress: handle.Argument(event, "mailAddress"),
			Validity:    time.Unix(int64(val), 0),
		},
		Signature: handle.Argument(event, "signature"),
	}
}

func main() {
	lambda.Start(handler)
}
