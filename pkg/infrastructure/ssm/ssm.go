package ssm

import (
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
	"github.com/ethereum/go-ethereum/log"
)

type (
	// Client ssm client
	Client struct {
		svc ssmiface.SSMAPI
	}
	// Key key
	Key string
)

const (
	keyPrefix              = "kaleido-"
	signingSecret          = keyPrefix + "signing-secret"
	slackToken             = keyPrefix + "slack-bot-token"
	emailSigningPublicKey  = keyPrefix + "email-signing-public-key"
	emailSigningPrivateKey = keyPrefix + "email-signing-private-key"
)

// New New client
func New() Client {
	svc := ssm.New(session.New(&aws.Config{Region: aws.String("us-east-1")}))
	//xray.AWS(svc.Client)
	return Client{
		svc: svc,
	}
}

// SigningSecret signing secret
func (c Client) SigningSecret(ctx context.Context) ([]byte, error) {
	k, err := c.get(ctx, withEnvSuffix(signingSecret))
	if err != nil {
		log.Error("get value failed", err)
		return []byte{}, err
	}
	return []byte(k), nil
}

// SlackToken slack token
func (c Client) SlackToken(ctx context.Context) (string, error) {
	return c.get(ctx, slackToken)
}

// EmailSigningPublicKey pub-key for email verification
func (c Client) EmailSigningPublicKey(ctx context.Context) (val ed25519.PublicKey, err error) {
	return c.getKey(ctx, withEnvSuffix(emailSigningPublicKey))
}

// EmailSigningPrivateKey private-key for email verification
func (c Client) EmailSigningPrivateKey(ctx context.Context) (val ed25519.PrivateKey, err error) {
	return c.getKey(ctx, withEnvSuffix(emailSigningPrivateKey))
}

func (c Client) getKey(ctx context.Context, key string) ([]byte, error) {
	k, err := c.get(ctx, key)
	if err != nil {
		return []byte{}, err
	}
	return hex.DecodeString(k)
}

// Get get parameter
func (c Client) get(ctx context.Context, key string) (val string, err error) {
	out, err := c.svc.GetParameterWithContext(ctx, &ssm.GetParameterInput{
		Name:           aws.String(key),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		log.Error("ssm get parameter failed", err)
		return
	}

	return *out.Parameter.Value, nil
}

func withEnvSuffix(key string) string {
	env := os.Getenv("EnvironmentId")
	if env == "prod" {
		return key
	}
	return key + "-" + "v1dev"
}
