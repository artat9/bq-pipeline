package auth

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	request "github.com/golang-jwt/jwt/request"
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
	secret, err := sec.Secret(ctx)
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
	req := http.Request{
		Header: http.Header{
			"Authorization": []string{val},
		},
	}
	t, err := request.ParseFromRequest(&req, request.AuthorizationHeaderExtractor, keyResolver(v.secret))
	if err != nil {
		return err
	}
	if !t.Valid {
		return fmt.Errorf("token invalid")
	}
	claims := t.Claims.(jwt.MapClaims)
	timestamp, err := strconv.ParseInt(claims["exp"].(string), 10, 64)
	if err != nil {
		return err
	}
	tm := time.Unix(timestamp, 0)
	if v.expired(tm) {
		return fmt.Errorf("token expired")
	}
	return nil
}

func (v Verifier) expired(t time.Time) bool {
	return v.cl.Now().Before(t)
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
