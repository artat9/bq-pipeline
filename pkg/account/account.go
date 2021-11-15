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
		Address     common.Address `json:"address"`
		Nonce       *big.Int       `json:"nonce"`
		Name        string         `json:"name"`
		Email       string         `json:"email"`
		Description string         `json:"description"`
	}

	// Service account service
	Service struct {
		rep Repository
	}
	// AddressVerifier verifier
	AddressVerifier interface {
		Recover(msg, sig string) (common.Address, error)
	}
	// Repository repository
	Repository interface {
		FromAddress(ctx context.Context, eoa common.Address) (Account, error)
		New(ctx context.Context, ac Account) error
		Update(ctx context.Context, ac Account) error
	}
	// SignService service
	SignService struct {
		verifier AddressVerifier
		rep      Repository
		signer   Signer
	}

	// UpdateInput update input
	UpdateInput struct {
		// TODO: updateable?
		//ID          *string `json:"email"`
		Name        *string `json:"name"`
		Email       *string `json:"email"`
		Description *string `json:"description"`
	}

	// VerifyService verify service
	VerifyService struct {
		verifier Verifier
	}

	// SignInInput input for signIn
	SignInInput struct {
		Msg string `json:"validity" validate:"required"`
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
		Verify(val string) (common.Address, error)
		Reflesh(refleshToken string) (string, error)
	}
)

// NewService new service
func NewService(r Repository) Service {
	return Service{
		rep: r,
	}
}

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

// Verify verify access token and return eoa
func (v VerifyService) Verify(ctx context.Context, accessToken string) (common.Address, error) {
	return v.verifier.Verify(accessToken)
}

// Update udpate account
func (s Service) Update(ctx context.Context, account common.Address, in UpdateInput) (Account, error) {
	ac, err := s.rep.FromAddress(ctx, account)
	if err != nil {
		return Account{}, err
	}
	if in.Description != nil {
		ac.Description = *in.Description
	}
	if in.Email != nil {
		ac.Email = *in.Email
	}
	if in.Name != nil {
		ac.Name = *in.Name
	}
	return ac, s.rep.Update(ctx, ac)
}

// Sign sign in / sign up
func (s SignService) Sign(ctx context.Context, in SignInInput) (SignInOutput, error) {
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

func (s SignService) signIn(ctx context.Context, non string, ac Account) (SignInOutput, error) {
	n := new(big.Int)
	n, _ = n.SetString(non, 10)
	if n.Cmp(ac.Nonce) != 0 {
		return SignInOutput{}, errors.New("invalid request")
	}
	at, rt, err := s.signer.Sign(ac.Address)
	if err != nil {
		return SignInOutput{}, err
	}
	newnonce, _ := nonce()
	ac.Nonce = newnonce
	if err = s.rep.Update(ctx, ac); err != nil {
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
