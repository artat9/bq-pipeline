package bigquery

import (
	"bq-pipeline/pkg/common/log"
	"bq-pipeline/pkg/user"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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
		Run(context.Context) (*bigquery.Job, error)
	}

	SampleReader interface {
		Read(ctx context.Context) (*bigquery.RowIterator, error)
	}
)

// Bigquery の Client を初期化
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

func (s Service) Upload(ctx context.Context) error {
	job, err := s.si.Run(ctx)

	if err != nil {
		return err
	}
	status, err := job.Wait(ctx)
	if err != nil {
		return err
	}
	if err := status.Err(); err != nil {
		return err
	}
	return nil
}

// func (s Service) Run(ctx context.Context){
// 	return s.si.Run(ctx)
// }

// func setCredentialEnv(string ){
// 	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/tmp/account.json")
// }

// func (s Service) Put(ctx context.Context, src interface{}) error {
// 	return nil
// }

//func (s Service) Insert()

// FIX 分離の実施
// func (s Service) Upload(ctx context.Context, client *bigquery.Client, u user.User) error {
// TODO FIX Table 名は環境変数で別出ししなくて大丈夫か
// TODO Streaming Insert でOK？それ以外だと読み込みジョブ方式がある
// memo
// source := bigquery.NewReaderSource(f)
// source.AutoDetect = true
// loader := client.Dataset(datasetID).Table(tableID).LoaderFrom(source)
// WriteTruncate（洗い替え）で書き込みする
// loader.LoadConfig.WriteDisposition = bigquery.WriteTruncate
// return client.Dataset(os.Getenv("TARGET_DATASET_ID")).Table("sample_terraform_user_table").Inserter().Put(ctx, u)
// return client.Dataset(os.Getenv("TARGET_DATASET_ID")).Table("sample_terraform_user_table").Inserter().Put(ctx, u)
// }

//func NewClient(ctx context.Context)(bigquery.Client){
//	return bigquery.NewClient(ctx, os.Getenv("TARGET_GCP_PROJECT_ID"))
//}

// TODO FIX
func NewSampleInserter(ctx context.Context, input []user.User, client *bigquery.Client) SampleInserter {
	var buf bytes.Buffer
	// struct から io.reader に変換
	err := json.NewEncoder(&buf).Encode(input)
	if err != nil {
		return nil
	}
	source := bigquery.NewReaderSource(&buf)
	fmt.Println("test", os.Getenv("TARGET_DATASET_ID"))

	loader := client.Dataset(os.Getenv("TARGET_DATASET_ID")).Table("sample_terraform_user_table").LoaderFrom(source)
	loader.WriteDisposition = bigquery.WriteAppend

	return loader
}

func NewSampleReader(ctx context.Context, client *bigquery.Client) SampleReader {
	query := "SELECT " +
		"id, name " +
		"FROM " +
		os.Getenv("TARGET_GCP_PROJECT_ID") +
		"." + os.Getenv("TARGET_DATASET_ID") +
		".sample_terraform_user_table"
	return client.Query(query)
}

// TODO DELETE
func (s Service) Read(ctx context.Context) (*bigquery.RowIterator, error) {
	return s.sr.Read(ctx)
}

// func (s Service) Download(ctx context.Context, it *bigquery.RowIterator) ([]*user.User, error) {

func (s *Service) Download(ctx context.Context) ([]*user.User, error) {
	it, err := s.sr.Read(ctx)
	if err != nil {
		log.Error("failed to read data", err)
		return nil, err
	}

	var outputData []*user.User
	for {
		var user user.User

		err = it.Next(&user)
		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Error("Failed to Iterate Query:%v", err)
			return outputData, err
		}

		log.Info(user)
		outputData = append(outputData, &user)
	}
	return outputData, nil
}
