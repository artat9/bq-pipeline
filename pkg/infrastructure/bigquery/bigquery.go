package bigquery

import (
	"bq-pipeline/pkg/common/log"
	"bq-pipeline/pkg/user"
	"context"
	"os"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type (
	Service struct {
		si SampleInserter
		sr SampleReader
	}
	GcpServiceAccountCredentialResolver interface {
		GcpServiceAccountCredential(ctx context.Context) (string, error)
	}
	SampleInserter interface {
		Put(ctx context.Context, src interface{}) error
	}

	SampleReader interface {
		Read(ctx context.Context) (*bigquery.RowIterator, error)
	}
)

const sample_user_table_1 = "sample_terraform_user_table_1"

func NewClient(ctx context.Context, sr GcpServiceAccountCredentialResolver) (*bigquery.Client, error) {
	c, err := sr.GcpServiceAccountCredential(ctx)
	if err != nil {
		log.Error("credential not found", err)
		return nil, err
	}
	data := []byte(c)
	return bigquery.NewClient(ctx, os.Getenv("TARGET_GCP_PROJECT_ID"), option.WithCredentialsJSON(data))
}

func NewService(si SampleInserter, sr SampleReader) Service {
	return Service{si: si, sr: sr}
}

func (s Service) Upload(ctx context.Context, users []user.User) error {
	if err := s.si.Put(ctx, users); err != nil {
		return err
	}
	return nil
}

func NewSampleInserter(client *bigquery.Client) SampleInserter {
	return client.Dataset(os.Getenv("TARGET_DATASET_ID")).Table(sample_user_table_1).Inserter()
}

func NewSampleReader(ctx context.Context, client *bigquery.Client) SampleReader {
	query := "SELECT " +
		"id, name " +
		"FROM " +
		os.Getenv("TARGET_GCP_PROJECT_ID") +
		"." + os.Getenv("TARGET_DATASET_ID") +
		"." + sample_user_table_1
	return client.Query(query)
}

func (s *Service) Download(ctx context.Context) ([]user.User, error) {
	it, err := s.sr.Read(ctx)
	if err != nil {
		log.Error("failed to read data", err)
		return nil, err
	}
	if it == nil {
		log.Info("read data is empty")
		return nil, nil
	}
	var outputData []user.User
	for {
		var user user.User

		err = it.Next(&user)
		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Error("failed to iterate query:%v", err)
			return outputData, err
		}
		outputData = append(outputData, user)
	}
	return outputData, nil
}
