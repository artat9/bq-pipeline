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
	// SignService service
	SignService struct {
		verifier AddressVerifier
		rep      Repository
		signer   Signer
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
	// Signer signer
	Signer interface {
		Sign(ad common.Address) (accessToken, refleshToken string, err error)
	}

	// Verifier verifier
	Verifier interface {
		Verify(val string) error
	}
)

// NewSignService new servcice
func NewSignService(v AddressVerifier, r Repository) SignService {
	return SignService{
		verifier: v,
		rep:      r,
	}
}

// SignIn sign in
func (s SignService) SignIn(ctx context.Context, msg, sig string) (SignInOutput, error) {
	ad, err := s.verifier.Recover(msg, sig)
	if err != nil {
		return SignInOutput{}, err
	}
	ac, err := s.rep.FromAddress(ctx, ad)
	if err != nil {
		return SignInOutput{}, err
	}
	if reflect.DeepEqual(ac, Account{}) {
		return s.signUp(ctx, ad)
	}
	return s.signIn(ctx, ac)
}

func (s SignService) signIn(ctx context.Context, ac Account) (SignInOutput, error) {
	return SignInOutput{}, nil
}

func (s SignService) signUp(ctx context.Context, address common.Address) (SignInOutput, error) {
	at, rt, err := s.signer.Sign(address)
	if err != nil {
		return SignInOutput{}, err
	}
	ac := New(address)
	return SignInOutput{
		Account:      ac,
		AccessToken:  at,
		RefleshToken: rt,
	}, s.rep.New(ctx, ac)
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
