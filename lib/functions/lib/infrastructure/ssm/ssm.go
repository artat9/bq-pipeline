package ssm

import (
	"context"
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
	keyPrefix     = "kaleido-"
	signingSecret = keyPrefix + "signing-secret"
)

// New New client
func New() Client {
	return Client{
		svc: ssm.New(session.New(&aws.Config{Region: aws.String("us-east-1")})),
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
	return key + "-" + "dev"
}
