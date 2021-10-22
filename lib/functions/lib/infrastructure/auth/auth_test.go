package auth

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type (
	fakeResolver struct {
		recoverAddressFunc func(msg, sig string) (common.Address, error)
	}
	fixedClock struct{}
)

func (f fixedClock) Now() time.Time {
	return time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC)
}

func (fr fakeResolver) RecoverAddress(msg, sig string) (common.Address, error) {
	return fr.recoverAddressFunc(msg, sig)
}

func fixedFakeResolver() fakeResolver {
	return fakeResolver{
		recoverAddressFunc: func(msg, sig string) (common.Address, error) {
			return common.HexToAddress("0x668F22f015BF2c91Bf4fb19a03619B8Ff593E8A8"), nil
		},
	}
}

func errorResolver() fakeResolver {
	return fakeResolver{
		recoverAddressFunc: func(msg, sig string) (common.Address, error) {
			return common.Address{}, errors.New("error")
		},
	}
}

func TestVerifier_Verify(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name    string
		v       Verifier
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Verifier{}
			if err := v.Verify(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Verifier.Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSigner_Sign(t *testing.T) {
	type fields struct {
		resolver Resolver
	}
	type args struct {
		msg string
		sig string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "new token should be issued",
			fields: fields{
				resolver: fixedFakeResolver(),
			},
			args: args{
				msg: "fake",
				sig: "fake",
			},
			want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZGRyZXNzIjoiMHg2NjhGMjJmMDE1QkYyYzkxQmY0ZmIxOWEwMzYxOUI4RmY1OTNFOEE4IiwiZXhwIjoxNjA5NDY0NjYxfQ.W4_mj-unUXcX_Ctkn4i1mhaSZfyoSymG590xpQjOBj8",
			wantErr: false,
		},
		{
			name: "new token should not be issued with failed eoa recovery",
			fields: fields{
				resolver: errorResolver(),
			},
			args: args{
				msg: "fake",
				sig: "fake",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Signer{
				resolver: tt.fields.resolver,
				secret:   []byte("secret"),
				cl:       fixedClock{},
			}
			got, err := s.Sign(tt.args.msg, tt.args.sig)
			if (err != nil) != tt.wantErr {
				t.Errorf("Signer.Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Signer.Sign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fakeResolver_RecoverAddress(t *testing.T) {
	type args struct {
		msg string
		sig string
	}
	tests := []struct {
		name    string
		fr      fakeResolver
		args    args
		want    common.Address
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fr.RecoverAddress(tt.args.msg, tt.args.sig)
			if (err != nil) != tt.wantErr {
				t.Errorf("fakeResolver.RecoverAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fakeResolver.RecoverAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSigner(t *testing.T) {
	type args struct {
		re     Resolver
		secret []byte
	}
	tests := []struct {
		name string
		args args
		want Signer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSigner(tt.args.re, tt.args.secret); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSigner() = %v, want %v", got, tt.want)
			}
		})
	}
}
