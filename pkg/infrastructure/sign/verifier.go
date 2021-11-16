package signer

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"kaleido-backend/pkg/common/log"
	"kaleido-backend/pkg/email"
	"time"
)

type (
	Verifier struct {
		publicKey ed25519.PublicKey
	}
	PublicKeyResolver interface {
		EmailSigningPublicKey(ctx context.Context) (val ed25519.PublicKey, err error)
	}
)

func NewVerifier(publicKey ed25519.PublicKey) Verifier {
	return Verifier{
		publicKey: publicKey,
	}
}

func NewVerifierWithResolver(ctx context.Context, r PublicKeyResolver) (Verifier, error) {
	publicKey, err := r.EmailSigningPublicKey(ctx)
	if err != nil {
		return Verifier{}, err
	}
	return NewVerifier(publicKey), err
}

func (v Verifier) Verify(in email.VerificationInput) bool {
	verified := verify(in, v.publicKey)
	if !verified {
		log.Info("sign not verified:" + in.Signature)
		return verified
	}
	return in.Validity.Before(time.Now())
}

func verify(in email.VerificationInput, key ed25519.PublicKey) bool {
	sig, err := hex.DecodeString(in.Signature)
	if err != nil {
		return false
	}
	if len(sig) != ed25519.SignatureSize {
		return false
	}
	var msg bytes.Buffer
	msg.WriteString(in.SignatureInput.String())
	return ed25519.Verify(key, msg.Bytes(), sig)
}
