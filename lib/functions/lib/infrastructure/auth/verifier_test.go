package auth

import (
	"context"
	"testing"
)

const (
	expired   = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZGRyZXNzIjoiMHg2NjhGMjJmMDE1QkYyYzkxQmY0ZmIxOWEwMzYxOUI4RmY1OTNFOEE4IiwiZXhwIjoxNTc3ODQyMjYxfQ.NzxbaqDHteGj0N_2UahzV8yVLGjr0sf9jgd2laF9RFI"
	falsified = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZGRyZXNzIjoiMHg2NjhGMjJmMDE1QkYyYzkxQmY0ZmIxOWEwMzYxOUI4RmY1OTNFOEE4IiwiZXhwIjoxNjA5NDY0NjYyfQ.7SeDFPR_V5WXOxobcW7XFJaAv6GTWNN2WxJgyNgJCMs"
)

func TestNewVerifier(t *testing.T) {
	type args struct {
		ctx context.Context
		sec SecretResolver
	}
	tests := []struct {
		name    string
		args    args
		want    Verifier
		wantErr bool
	}{
		{
			name:    "initialization success",
			args:    args{ctx: context.Background(), sec: fixedFakeSecretResolver()},
			wantErr: false,
		},
		{
			name:    "fail in case get secret failed",
			args:    args{ctx: context.Background(), sec: errorSecretResolver()},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewVerifier(tt.args.ctx, tt.args.sec)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSigner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestVerifier_Verify(t *testing.T) {
	secret, _ := fixedFakeSecretResolver().Secret(context.TODO())
	type fields struct {
		secret []byte
		cl     Clock
	}
	type args struct {
		val string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success verification",
			fields: fields{
				secret: secret,
				cl:     fixedClock{},
			},
			args: args{
				val: accessToken,
			},
		},
		{
			name: "error with an invalid token",
			fields: fields{
				secret: secret,
				cl:     fixedClock{},
			},
			args: args{
				val: "",
			},
			wantErr: true,
		},
		{
			name: "error with an expired token",
			fields: fields{
				secret: secret,
				cl:     fixedClock{},
			},
			args: args{
				val: expired,
			},
			wantErr: true,
		},
		{
			name: "error with a falsified token",
			fields: fields{
				secret: secret,
				cl:     fixedClock{},
			},
			args: args{
				val: falsified,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Verifier{
				secret: secret,
				cl:     tt.fields.cl,
			}
			if err := v.Verify(tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("Verifier.Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
