package email

import (
	"context"
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

	NotificationService struct {
		signer   Signer
		notifier Notifier
	}

	VerificationService struct {
		verifier Verifier
	}

	// Signer signer
	Signer interface {
		Sign(mailAddress string) VerificationInput
	}
	// Verifier verifyer
	Verifier interface {
		Verify(input VerificationInput) (bool, error)
	}
	// Notifier notifier
	Notifier interface {
		SendApplicationCreated(ctx context.Context, in VerificationInput) error
	}
)

func NewNotificationService(signer Signer, notifier Notifier) NotificationService {
	return NotificationService{
		signer:   signer,
		notifier: notifier,
	}
}

// String string
func (in SignatureInput) String() string {
	return fmt.Sprintf("mailAddress=%s&validity=%d", in.MailAddress, in.Validity.UnixNano())
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
func (s NotificationService) SendEmailVerification(ctx context.Context, email string) error {
	return s.notifier.SendApplicationCreated(ctx, s.signer.Sign(email))
}

// VerifyEmail verify email
func (s VerificationService) VerifyEmail(in VerificationInput) (bool, error) {
	return s.verifier.Verify(in)
}
