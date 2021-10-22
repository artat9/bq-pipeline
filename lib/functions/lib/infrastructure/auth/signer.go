package auth

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/golang-jwt/jwt/v4"
)

const (
	accessTokenValidity  = time.Minute * 30
	refleshTokenValidity = time.Hour * 24 * 14
)

type (

	// Signer token signer
	Signer struct {
		secret []byte
		cl     Clock
	}
	// Clock clock
	Clock interface {
		Now() time.Time
	}
	// SecretResolver resolver for secret
	SecretResolver interface {
		SigningSecret(ctx context.Context) ([]byte, error)
	}
	realClock struct{}
)

func (r realClock) Now() time.Time { return time.Now() }

// NewSigner new signer
func NewSigner(ctx context.Context, sec SecretResolver) (Signer, error) {
	secret, err := sec.SigningSecret(ctx)
	if err != nil {
		return Signer{}, err
	}
	return Signer{
		secret: secret,
		cl:     realClock{},
	}, err
}

// Sign sign
func (s Signer) Sign(ad common.Address) (accessToken, refleshToken string, err error) {
	accessToken, err = s.newAccessToken(ad)
	if err != nil {
		return
	}
	refleshToken, err = s.newRefleshToken(ad)
	if err != nil {
		return
	}
	return
}

func (s Signer) newAccessToken(ad common.Address) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["address"] = ad.Hex()
	claims["exp"] = jwt.NewNumericDate(s.cl.Now().Add(accessTokenValidity))
	t, err := token.SignedString(s.secret)
	return t, err
}

func newAccessToken(ad common.Address, secret []byte, now time.Time) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["address"] = ad.Hex()
	claims["exp"] = jwt.NewNumericDate(now.Add(accessTokenValidity))
	t, err := token.SignedString(secret)
	return t, err
}

func (s Signer) newRefleshToken(ad common.Address) (string, error) {
	return newRefleshToken(ad, s.secret, s.cl.Now())
}

func newRefleshToken(ad common.Address, secret []byte, now time.Time) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = ad.Hex()
	claims["exp"] = jwt.NewNumericDate(now.Add(refleshTokenValidity))
	t, err := token.SignedString(secret)
	return t, err
}
