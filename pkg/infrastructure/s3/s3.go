package s3

import (
	"bq-pipeline/pkg/common/log"
	"context"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// S3 s3
type S3 struct {
	svc s3iface.S3API
}

// New new s3
func New() S3 {
	cli := s3.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))
	//xray.AWS(cli.Client)
	return S3{
		svc: cli,
	}
}

// IssueUploadURL issue upload url
func (s S3) IssueUploadURL(ctx context.Context, filename string) (url string, err error) {
	req, _ := s.svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(resolveBucketName()),
		Key:    aws.String(filename),
	})
	url, err = req.Presign(5 * time.Minute)
	if err != nil {
		log.Error("presign failed", err)
	}
	return
}

func resolveBucketName() string {
	return "kaleido-backend-asset-" + os.Getenv("EnvironmentId")
}
