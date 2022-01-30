package bigquery

import (
	"bq-pipeline/pkg/common/log"
	"bq-pipeline/pkg/user"
	"context"
	"os"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

type (
   Service struct {}
	 BQ interface{
		 New(ctx context.Context) (*bigquery.Client, error)
		 Upload(ctx context.Context, client bigquery.Client,user user.User) error
		 Download(ctx context.Context, client bigquery.Client)
	 }

)

func New(ctx context.Context) (*bigquery.Client, error) {
	return bigquery.NewClient(ctx, os.Getenv("TARGET_GCP_PROJECT_ID"))
}

func (s Service) Upload(ctx context.Context, client *bigquery.Client, u user.User) error {
	// TODO FIX Table 名は環境変数で別出ししなくて大丈夫か
	return client.Dataset(os.Getenv("TARGET_DATASET_ID")).Table("sample_terraform_user_table").Inserter().Put(ctx, u)
}

func (s Service) Download(ctx context.Context, client bigquery.Client, query string)([]*user.User, error) {

	it, err := client.Query(query).Read(ctx)

	if err != nil {
    log.Error("Failed to Read Query:%v", err)
  }

	var outputData[]*user.User
	for {
		var user user.User

    err := it.Next(&user)
    if err == iterator.Done {
        break
    }

    if err != nil {
        log.Error("Failed to Iterate Query:%v", err)
				return outputData, nil
    }

		log.Info(user)
		outputData = append(outputData, &user)
   }
	 return outputData, nil
}
