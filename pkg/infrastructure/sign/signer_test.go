package signer

import (
	"context"
	"kaleido-backend/pkg/infrastructure/ssm"
	"os"
	"testing"
)

func TestSign(t *testing.T) {
	os.Setenv("EnvironmentId", "v1dev")
	signer, _ := NewWithResolver(context.Background(), ssm.New())
	in := signer.Sign("yushi.masui@bridges.inc")
	verifier, _ := NewVerifierWithResolver(context.Background(), ssm.New())
	if !verifier.Verify(in) {
		t.Error()
	}
}
