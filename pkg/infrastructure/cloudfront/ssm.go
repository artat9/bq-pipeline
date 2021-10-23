package cloudfront

import (
	"context"
	"kaleido-backend/pkg/common/log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/aws/aws-sdk-go/service/cloudfront/cloudfrontiface"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/google/uuid"
)

type (
	// Client ssm client
	Client struct {
		svc cloudfrontiface.CloudFrontAPI
	}
)

// New new client
func New() Client {
	svc := cloudfront.New(session.New(&aws.Config{Region: aws.String("us-east-1")}))
	xray.AWS(svc.Client)
	return Client{
		svc: svc,
	}
}

// Invalidate invalidate cache for items
func (c Client) Invalidate(ctx context.Context, items []string) error {
	tar := []string{}
	for _, i := range items {
		if strings.HasPrefix(i, "/") {
			tar = append(tar, i)
		} else {
			tar = append(tar, "/"+i)
		}
	}
	if _, err := c.svc.CreateInvalidationWithContext(ctx, &cloudfront.CreateInvalidationInput{
		DistributionId: aws.String(os.Getenv("AssetDistributionId")),
		InvalidationBatch: &cloudfront.InvalidationBatch{
			CallerReference: aws.String(uuid.NewString()),
			Paths:           &cloudfront.Paths{Quantity: aws.Int64(int64(len(tar))), Items: aws.StringSlice(tar)},
		},
	}); err != nil {
		log.Error("create invalidation failed", err)
		return err
	}
	return nil
}
