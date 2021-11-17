package email

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type (

	// SignatureInput confirmation
	SignatureInput struct {
		MailAddress string    `json:"mailAddress"`
		Validity    time.Time `json:"validity"`
	}
	// VerificationInput verification input
	VerificationInput struct {
		SignatureInput
		Signature string `json:"signature"`
	}

	SignService struct {
		signer   Signer
		notifier Notifier
	}

	VerificationService struct {
		verifier Verifier
		notifier InhouseNotifier
	}

	// Signer signer
	Signer interface {
		Sign(mailAddress string) VerificationInput
	}
	// Verifier verifyer
	Verifier interface {
		Verify(input VerificationInput) bool
	}
	// Notifier notifier
	Notifier interface {
		SendApplicationCreated(ctx context.Context, in VerificationInput) error
	}
	InhouseNotifier interface {
		NotifyEmailAddressConfirmationCompleted(ctx context.Context, in VerificationInput) error
	}
)

func NewSignService(signer Signer, notifier Notifier) SignService {
	return SignService{
		signer:   signer,
		notifier: notifier,
	}
}

func NewVerificationService(verifier Verifier, not InhouseNotifier) VerificationService {
	return VerificationService{
		verifier: verifier,
		notifier: not,
	}
}

// String string
func (in SignatureInput) String() string {
	return fmt.Sprintf("mailAddress=%s&validity=%d", in.MailAddress, in.Validity.Unix())
}

// QueryString query string
func (veri VerificationInput) QueryString() string {
	return veri.String() + "&signature=" + veri.Signature
}

// NewInput new signature input
func NewInput(mailAddress string) SignatureInput {
	return SignatureInput{
		MailAddress: mailAddress,
		Validity:    time.Now().Add(time.Hour),
	}
}

// SendEmailVerification send email verification
func (s SignService) SendEmailVerification(ctx context.Context, email string) error {
	return s.notifier.SendApplicationCreated(ctx, s.signer.Sign(email))
}

// VerifyEmail verify email
func (s VerificationService) VerifyEmail(ctx context.Context, in VerificationInput) (string, error) {
	verified := s.verifier.Verify(in)
	if !verified {
		return "", errors.New("unauthorized")
	}
	return in.MailAddress, s.notifier.NotifyEmailAddressConfirmationCompleted(ctx, in)
}
