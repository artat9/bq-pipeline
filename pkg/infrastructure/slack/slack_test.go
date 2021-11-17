package slack

import (
	"context"
	"kaleido-backend/pkg/email"
	"kaleido-backend/pkg/infrastructure/ssm"
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
