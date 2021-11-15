package account

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

type (
	fakeRep struct {
		FromAddressFunc func(ctx context.Context, eoa common.Address) (Account, error)
		NewFunc         func(ctx context.Context, ac Account) error
		UpdateFunc      func(ctx context.Context, ac Account) error
	}
)

func (fr fakeRep) FromAddress(ctx context.Context, eoa common.Address) (Account, error) {
	return fr.FromAddressFunc(ctx, eoa)
}
func (fr fakeRep) New(ctx context.Context, ac Account) error    { return fr.NewFunc(ctx, ac) }
func (fr fakeRep) Update(ctx context.Context, ac Account) error { return fr.UpdateFunc(ctx, ac) }
func newFakeRep() fakeRep {
	return fakeRep{
		FromAddressFunc: func(ctx context.Context, eoa common.Address) (Account, error) { return Account{Address: eoa}, nil },
		NewFunc:         func(ctx context.Context, ac Account) error { return nil },
		UpdateFunc:      func(ctx context.Context, ac Account) error { return nil },
	}
}
func newErrorFakeRep() fakeRep {
	return fakeRep{
		FromAddressFunc: func(ctx context.Context, eoa common.Address) (Account, error) { return Account{}, errors.New("err") },
		NewFunc:         func(ctx context.Context, ac Account) error { return errors.New("err") },
		UpdateFunc:      func(ctx context.Context, ac Account) error { return errors.New("err") },
	}
}

func TestNew(t *testing.T) {
	addr := common.HexToAddress("0x668F22f015BF2c91Bf4fb19a03619B8Ff593E8A8")
	u1 := New(addr)
	u2 := New(addr)
	t.Run("address should be matched", func(t *testing.T) {
		if addr != u1.Address {
			t.Error("got:", u1.Address)
		}
	})
	t.Run("nonce should be a random value", func(t *testing.T) {
		if u1.Nonce.Cmp(u2.Nonce) == 0 {
			t.Error("got:", u1.Nonce)
		}
	})
}

func TestNewService(t *testing.T) {
	type args struct {
		r Repository
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerifyService_Verify(t *testing.T) {
	type fields struct {
		verifier Verifier
	}
	type args struct {
		ctx         context.Context
		accessToken string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    common.Address
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := VerifyService{
				verifier: tt.fields.verifier,
			}
			got, err := v.Verify(tt.args.ctx, tt.args.accessToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyService.Verify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VerifyService.Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Update(t *testing.T) {
	rep := newFakeRep()
	exp := "test"
	type fields struct {
		rep Repository
	}
	type args struct {
		ctx     context.Context
		account common.Address
		in      UpdateInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Account
		wantErr bool
	}{
		{
			name:    "name should be updated",
			fields:  fields{rep: rep},
			args:    args{ctx: context.Background(), account: common.Address{}, in: UpdateInput{Name: &exp}},
			want:    Account{Name: exp},
			wantErr: false,
		},
		{
			name:    "email should be updated",
			fields:  fields{rep: rep},
			args:    args{ctx: context.Background(), account: common.Address{}, in: UpdateInput{Email: &exp}},
			want:    Account{Email: exp},
			wantErr: false,
		},
		{
			name:    "description should be updated",
			fields:  fields{rep: rep},
			args:    args{ctx: context.Background(), account: common.Address{}, in: UpdateInput{Description: &exp}},
			want:    Account{Description: exp},
			wantErr: false,
		},
		{
			name:    "handle rep error",
			fields:  fields{rep: newErrorFakeRep()},
			args:    args{ctx: context.Background(), account: common.Address{}, in: UpdateInput{Description: &exp}},
			want:    Account{},
			wantErr: true,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				rep: tt.fields.rep,
			}
			got, err := s.Update(tt.args.ctx, tt.args.account, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
