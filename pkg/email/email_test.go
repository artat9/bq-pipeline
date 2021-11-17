package email

import (
	"testing"
	"time"
)

func TestQueryString(t *testing.T) {
	t.Run("query string should be expressed", func(t *testing.T) {
		input := VerificationInput{
			SignatureInput: SignatureInput{
				MailAddress: "test@example.com",
				Validity:    time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC),
			},
			Signature: "test",
		}

		got := input.QueryString()
		want := "mailAddress=test@example.com&validity=946688461&signature=test"
		if want != got {
			t.Error("got ", got)
		}
	})
}
