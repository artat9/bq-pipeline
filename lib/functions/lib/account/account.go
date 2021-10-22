package account

import (
	"context"
	"crypto/rand"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
)

type (
	// Account account
	Account struct {
		Address common.Address `json:"address"`
		Nonce   *big.Int       `json:"nonce"`
	}
	// AddressVerifier verifier
	AddressVerifier interface {
		Recover(msg, sig string) (common.Address, error)
	}
	// Repository repository
	Repository interface {
		FromAddress(ctx context.Context, eoa common.Address) (Account, error)
		New(ctx context.Context, ac Account) error
	}
	// Service service
	Service struct {
		verifier AddressVerifier
		rep      Repository
	}

	// SignInInput input for signIn
	SignInInput struct {
		Msg string `json:"msg"`
		Sig string `json:"sig"`
	}

	// SignInOutput output for signin
	SignInOutput struct {
		Account
		AccessToken  string `json:"accessToken"`
		RefleshToken string `json:"refleshToken"`
	}
)

// NewService new servcice
func NewService(v AddressVerifier, r Repository) Service {
	return Service{
		verifier: v,
		rep:      r,
	}
}

// SignIn sign in
func (s Service) SignIn(ctx context.Context, msg, sig string) (Account, error) {
	ad, err := s.verifier.Recover(msg, sig)
	if err != nil {
		return Account{}, err
	}
	ac, err := s.rep.FromAddress(ctx, ad)
	if err != nil {
		return Account{}, err
	}
	if reflect.DeepEqual(ac, Account{}) {
		return s.signUp(ctx, ad)
	}
	return s.signIn(ctx, ac)
}

func (s Service) signIn(ctx context.Context, ac Account) (Account, error) {
	return ac, nil
}

func (s Service) signUp(ctx context.Context, address common.Address) (Account, error) {
	ac := New(address)
	return ac, s.rep.New(ctx, ac)
}

// New new account
func New(address common.Address) Account {
	n, _ := nonce()
	return Account{
		Address: address,
		Nonce:   n,
	}
}

func nonce() (n *big.Int, err error) {
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(130), nil).Sub(max, big.NewInt(1))
	return rand.Int(rand.Reader, max)
}
