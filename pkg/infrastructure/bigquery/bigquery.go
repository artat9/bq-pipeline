package bigquery

import (
	"bq-pipeline/pkg/user"
	"context"

	"cloud.google.com/go/bigquery"
)

type (
	Service struct {
	}
	BQ interface {
	}
)

func (s Service) Upload(ctx context.Context, u user.User) error {
	c, err := bigquery.NewClient(ctx, "my-project-ID")
	if err != nil {
		return err
	}
	return c.Dataset("Test").Table("test-table").Inserter().Put(ctx, u)

}
