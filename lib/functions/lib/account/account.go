package account

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/go-playground/validator.v9"
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

	// VerifyService verify service
	VerifyService struct {
		verifier Verifier
	}

	// SignInInput input for signIn
	SignInInput struct {
		Msg string `json:"msg" validate:"required"`
		Sig string `json:"sig" validate:"required"`
	}

	// RefleshInput reflesh
	RefleshInput struct {
		RefleshToken string `json:"refleshToken"`
	}

	// RefleshOutput reflesh
	RefleshOutput struct {
		AccessToken string `json:"accessToken"`
	}

	// SignInOutput output for signin
	SignInOutput struct {
		AccessToken  string         `json:"accessToken"`
		RefleshToken string         `json:"refleshToken"`
		Address      common.Address `json:"address"`
	}
	// Signer signer
	Signer interface {
		Sign(ad common.Address) (accessToken, refleshToken string, err error)
	}

	// Verifier verifier
	Verifier interface {
		Verify(val string) error
		Reflesh(refleshToken string) (string, error)
	}
)

// NewSignService new servcice
func NewSignService(v AddressVerifier, r Repository, s Signer) SignService {
	return SignService{
		verifier: v,
		rep:      r,
		signer:   s,
	}
}

// NewVerifyService new verify service
func NewVerifyService(v Verifier) VerifyService {
	return VerifyService{
		verifier: v,
	}
}

// Reflesh renew token
func (v VerifyService) Reflesh(ctx context.Context, in RefleshInput) (RefleshOutput, error) {
	out, err := v.verifier.Reflesh(in.RefleshToken)
	return RefleshOutput{
		AccessToken: out,
	}, err
}

// Sign sign in / sign up
func (s SignService) Sign(ctx context.Context, in SignInInput) (SignInOutput, error) {
	// TODO: verify nonce
	if err := validator.New().Struct(&in); err != nil {
		return SignInOutput{}, err
	}
	ad, err := s.verifier.Recover(in.Msg, in.Sig)
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
	return s.signIn(ctx, in.Msg, ac)
}

func (s SignService) signIn(ctx context.Context, nonce string, ac Account) (SignInOutput, error) {
	n := new(big.Int)
	n, _ = n.SetString(nonce, 10)
	if n.Cmp(ac.Nonce) != 0 {
		return SignInOutput{}, errors.New("invalid request")
	}
	at, rt, err := s.signer.Sign(ac.Address)
	if err != nil {
		return SignInOutput{}, err
	}
	if err = s.rep.New(ctx, New(ac.Address)); err != nil {
		return SignInOutput{}, err
	}
	return SignInOutput{
		Address:      ac.Address,
		AccessToken:  at,
		RefleshToken: rt,
	}, nil
}

func (s SignService) signUp(ctx context.Context, address common.Address) (SignInOutput, error) {
	at, rt, err := s.signer.Sign(address)
	if err != nil {
		return SignInOutput{}, err
	}
	ac := New(address)
	return SignInOutput{
		Address:      ac.Address,
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
