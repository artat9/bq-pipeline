package main

import (
	"context"
	"kaleido-backend/pkg/email"
	"kaleido-backend/pkg/infrastructure/ses"
	signer "kaleido-backend/pkg/infrastructure/sign"
	"kaleido-backend/pkg/infrastructure/ssm"
	"os"
)

func main() {
	os.Setenv("RootDomain", "kaleido-v1dev.tk")
	os.Setenv("EnvironmentId", "v1dev")
	signer, err := signer.NewWithResolver(context.Background(), ssm.New())
	if err != nil {
		panic(err)
	}
	email.NewNotificationService(signer, ses.New()).SendEmailVerification(context.Background(), "dev@bridges.inc")
}
