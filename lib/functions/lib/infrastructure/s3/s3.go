package s3

import (
	"aurora-backend/lib/functions/lib/common/log"
	"context"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// S3 s3
type S3 struct {
	svc s3iface.S3API
}

// New new s3
func New() S3 {
	cli := s3.New(session.New())
	return S3{
		svc: cli,
	}
}

// UploadPostTitleImage uplaod pj title image
func (s S3) UploadPostTitleImage(ctx context.Context, txID string, content io.ReadSeeker) error {
	return s.uploadImage(ctx, postTitleKey(txID), content)
}

// UploadPostDescriptionImage upload description image
func (s S3) UploadPostDescriptionImage(ctx context.Context, postID string, content io.ReadSeeker) error {
	return s.uploadImage(ctx, postDescriptionKey(postID), content)
}

// DownloadPostTitleImage download post title image
func (s S3) DownloadPostTitleImage(ctx context.Context, postID string) (io.ReadCloser, error) {
	return s.downloadImage(ctx, postTitleKey(postID))
}

// DownloadPostDescriptionImage download post description image
func (s S3) DownloadPostDescriptionImage(ctx context.Context, postID string) (io.ReadCloser, error) {
	return s.downloadImage(ctx, postDescriptionKey(postID))
}

func (s S3) downloadImage(ctx context.Context, key string) (io.ReadCloser, error) {
	return s.getObject(ctx, resolveBucketName(), key)
}

func (s S3) uploadImage(ctx context.Context, key string, content io.ReadSeeker) error {
	return s.putObject(ctx, resolveBucketName(), key, content)
}

func postDescriptionKey(postID string) string {
	return "projectTitle-" + postID
}
func postTitleKey(postID string) string {
	return "projectTitle-" + postID
}

func (s S3) putObject(ctx context.Context, bucket, key string, content io.ReadSeeker) error {
	in := &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        content,
		ContentType: aws.String("image/png"),
	}
	_, err := s.svc.PutObjectWithContext(ctx, in)
	if err != nil {
		log.Error("put object failed", err)
	}
	return err
}

func (s S3) getObject(ctx context.Context, bucket, key string) (io.ReadCloser, error) {
	in := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	res, err := s.svc.GetObjectWithContext(ctx, in)
	if err != nil {
		log.Error("get object failed", err)
	}
	return res.Body, err

}

func resolveBucketName() string {
	return os.Getenv("AssetBucketName")
}
