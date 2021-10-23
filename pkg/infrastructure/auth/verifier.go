package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/golang-jwt/jwt/v4"
	request "github.com/golang-jwt/jwt/v4/request"
)

type (
	// Verifier token verifier
	Verifier struct {
		secret []byte
		cl     Clock
	}
)

// NewVerifier new verifier
func NewVerifier(ctx context.Context, sec SecretResolver) (Verifier, error) {
	secret, err := sec.SigningSecret(ctx)
	if err != nil {
		return Verifier{}, err
	}
	return Verifier{
		secret: secret,
		cl:     realClock{},
	}, err
}

// Verify verify claim
func (v Verifier) Verify(val string) error {
	_, err := v.verify(val)
	return err
}

func (v Verifier) verify(val string) (jwt.MapClaims, error) {
	jwt.TimeFunc = v.cl.Now
	req := http.Request{
		Header: http.Header{
			"Authorization": []string{val},
		},
	}
	t, err := request.ParseFromRequest(&req, request.AuthorizationHeaderExtractor, keyResolver(v.secret))
	if err != nil {
		return jwt.MapClaims{}, err
	}
	if !t.Valid {
		return jwt.MapClaims{}, fmt.Errorf("token invalid")
	}
	return t.Claims.(jwt.MapClaims), nil
}

// Reflesh renew access token
func (v Verifier) Reflesh(refleshToken string) (string, error) {
	cls, err := v.verify(refleshToken)
	if err != nil {
		return "", err
	}
	ad, ok := cls["sub"]
	if !ok {
		return "", fmt.Errorf("token invalid")
	}
	return newAccessToken(common.HexToAddress(ad.(string)), v.secret, v.cl.Now())
}

func keyResolver(secret []byte) func(t *jwt.Token) (interface{}, error) {
	return func(t *jwt.Token) (interface{}, error) {
		_, err := t.Method.(*jwt.SigningMethodHMAC)
		if !err {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return secret, nil
	}
}
