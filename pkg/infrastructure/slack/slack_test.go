package slack

import (
	"bq-pipeline/pkg/email"
	"bq-pipeline/pkg/infrastructure/ssm"
	"context"
	"os"
	"testing"
	"time"
)

func TestNotifyEmailAddressConfirmationCompleted(t *testing.T) {
	os.Setenv("EnvironmentId", "v1dev")
	cli, _ := New(context.Background(), ssm.New())
	cli.NotifyEmailAddressConfirmationCompleted(context.Background(), email.VerificationInput{
		SignatureInput: email.SignatureInput{
			MailAddress: "test@example.com",
			Validity:    time.Now(),
		},
		Signature: "test",
	})
}
