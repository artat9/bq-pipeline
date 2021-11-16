package signer

import (
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"kaleido-backend/pkg/email"
)

type (
	Signer struct {
		privateKey ed25519.PrivateKey
	}
	PrivateKeyResolver interface {
		EmailSigningPrivateKey(ctx context.Context) (val ed25519.PrivateKey, err error)
	}
)

func New(privateKey ed25519.PrivateKey) Signer {
	return Signer{
		privateKey: privateKey,
	}
}

func NewWithResolver(ctx context.Context, r PrivateKeyResolver) (Signer, error) {
	pivateKey, err := r.EmailSigningPrivateKey(ctx)
	if err != nil {
		return Signer{}, err
	}
	return New(pivateKey), err
}

func (s Signer) Sign(mailAddress string) email.VerificationInput {
	in := email.NewInput(mailAddress)
	sign := hex.EncodeToString(ed25519.Sign(s.privateKey, []byte(in.String())))
	return email.VerificationInput{
		SignatureInput: in,
		Signature:      sign,
	}
}
