package auth

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

const (
	accessToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZGRyZXNzIjoiMHg2NjhGMjJmMDE1QkYyYzkxQmY0ZmIxOWEwMzYxOUI4RmY1OTNFOEE4IiwiZXhwIjoxNjA5NDY0NjYxfQ.W4_mj-unUXcX_Ctkn4i1mhaSZfyoSymG590xpQjOBj8"
)

type (
	fakeSecretResolver struct {
		resolveFunc func(ctx context.Context) ([]byte, error)
	}
	fixedClock struct{}
)

func (f fixedClock) Now() time.Time {
	return time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC)
}

func (fs fakeSecretResolver) SigningSecret(ctx context.Context) ([]byte, error) {
	return fs.resolveFunc(ctx)
}

func fixedFakeSecretResolver() fakeSecretResolver {
	return fakeSecretResolver{
		resolveFunc: func(ctx context.Context) ([]byte, error) {
			return []byte("secret"), nil
		},
	}
}

func errorSecretResolver() fakeSecretResolver {
	return fakeSecretResolver{
		resolveFunc: func(ctx context.Context) ([]byte, error) {
			return nil, errors.New("error")
		},
	}
}

func TestSigner_Sign(t *testing.T) {

	type args struct {
		msg string
		sig string
	}
	tests := []struct {
		name             string
		want             string
		refleshTokenWant string
		wantErr          bool
	}{
		{
			name: "new token should be issued",

			want:             accessToken,
			wantErr:          false,
			refleshTokenWant: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTA2NzI0NjEsInN1YiI6IjB4NjY4RjIyZjAxNUJGMmM5MUJmNGZiMTlhMDM2MTlCOEZmNTkzRThBOCJ9.V0C3rh5ceFOb2vkwsJYdlMUPUCHzqLD4gj3PEZ2khwY",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Signer{
				secret: []byte("secret"),
				cl:     fixedClock{},
			}
			got, refGot, err := s.Sign(common.HexToAddress("0x668F22f015BF2c91Bf4fb19a03619B8Ff593E8A8"))
			if (err != nil) != tt.wantErr {
				t.Errorf("Signer.Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Signer.Sign() = %v, want %v", got, tt.want)
			}
			if refGot != tt.refleshTokenWant {
				t.Errorf("Signer.Sign() = %v, want %v", refGot, tt.refleshTokenWant)
			}
		})
	}
}

func TestSigner_newRefleshToken(t *testing.T) {
	type fields struct {
		secret []byte
		cl     Clock
	}
	type args struct {
		ad common.Address
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Signer{
				secret: tt.fields.secret,
				cl:     tt.fields.cl,
			}
			got, err := s.newRefleshToken(tt.args.ad)
			if (err != nil) != tt.wantErr {
				t.Errorf("Signer.newRefleshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Signer.newRefleshToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_realClock_Now(t *testing.T) {
	tests := []struct {
		name string
		r    realClock
		want time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := realClock{}
			if got := r.Now(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("realClock.Now() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSigner(t *testing.T) {
	type args struct {
		ctx context.Context
		sec SecretResolver
	}
	tests := []struct {
		name    string
		args    args
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
			_, err := NewSigner(tt.args.ctx, tt.args.sec)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSigner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSigner_newAccessToken(t *testing.T) {
	type fields struct {
		secret []byte
		cl     Clock
	}
	type args struct {
		ad common.Address
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Signer{
				secret: tt.fields.secret,
				cl:     tt.fields.cl,
			}
			got, err := s.newAccessToken(tt.args.ad)
			if (err != nil) != tt.wantErr {
				t.Errorf("Signer.newAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Signer.newAccessToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
