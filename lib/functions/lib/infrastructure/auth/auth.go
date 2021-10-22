package auth

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/golang-jwt/jwt"
)

const (
	accessTokenValidity  = time.Minute * 30
	refleshTokenValidity = time.Hour * 24 * 14
)

type (
	// Verifier token verifier
	Verifier struct {
	}
	// Signer token signer
	Signer struct {
		resolver Resolver
		secret   []byte
		cl       Clock
	}
	// Resolver eoa resolver
	Resolver interface {
		RecoverAddress(msg, sig string) (common.Address, error)
	}
	// Clock clock
	Clock interface {
		Now() time.Time
	}
	realClock struct{}
)

func (r realClock) Now() time.Time { return time.Now() }

// Verify verify claim
func (v Verifier) Verify(msg string) error {
	return nil
}

// NewSigner new signer
func NewSigner(re Resolver, secret []byte) Signer {
	return Signer{
		resolver: re,
		secret:   secret,
		cl:       realClock{},
	}
}

// Sign sign
func (s Signer) Sign(msg, sig string) (string, error) {
	ad, err := s.resolver.RecoverAddress(msg, sig)
	if err != nil {
		return "", err
	}
	return s.newAccessToken(ad)
}

func (s Signer) newAccessToken(ad common.Address) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["address"] = ad.Hex()
	claims["exp"] = s.cl.Now().Add(accessTokenValidity).Unix()
	t, err := token.SignedString(s.secret)
	return t, err
}

func (s Signer) newRefleshToken(ad common.Address) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = ad.Hex()
	claims["exp"] = s.cl.Now().Add(refleshTokenValidity).Unix()
	t, err := token.SignedString(s.secret)
	return t, err
}
