package bigquery

import (
	"bq-pipeline/pkg/user"
	"context"

	"cloud.google.com/go/bigquery"
)

type (
	// BqService struct {
	  // bqService BQ
	// }
	// BQ interface {
		// Put(ctx context.Context, u user.User)
	// }
   BQ struct {
		 service Service
	 }

	 Service interface{
		 Put(ctx context.Context, u user.User)
	 }

)

func New(ctx context.Context) BQ {
  // return BqService {
	// 	bqService: bqService,
	// }
	return BQ{
		service: bigquery.NewClient(ctx, "my-project-ID"),
	}
}

func (s BqService) Put(ctx context.Context, u user.User) error {
	//c, err := bigquery.NewClient(ctx, "my-project-ID")
	if err != nil {
		return err
	}
	return c.Dataset("Test").Table("test-table").Inserter().Put(ctx, u)
}
