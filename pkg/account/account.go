package account

import (
	"context"
	"crypto/rand"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
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

// NewVerifyService new verify service
func NewVerifyService(v Verifier) VerifyService {
	return VerifyService{
		verifier: v,
	}
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
