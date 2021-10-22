package auth

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewVerifier(tt.args.ctx, tt.args.sec)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewVerifier() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVerifier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerifier_Verify(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Verifier{
				secret: tt.fields.secret,
				cl:     tt.fields.cl,
			}
			if err := v.Verify(tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("Verifier.Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVerifier_expired(t *testing.T) {
	type fields struct {
		secret []byte
		cl     Clock
	}
	type args struct {
		t time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Verifier{
				secret: tt.fields.secret,
				cl:     tt.fields.cl,
			}
			if got := v.expired(tt.args.t); got != tt.want {
				t.Errorf("Verifier.expired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_keyResolver(t *testing.T) {
	type args struct {
		secret []byte
	}
	tests := []struct {
		name string
		args args
		want func(t *jwt.Token) (interface{}, error)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := keyResolver(tt.args.secret); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keyResolver() = %v, want %v", got, tt.want)
			}
		})
	}
}
