package bigquery

import (
	"bq-pipeline/pkg/infrastructure/ssm"
	"context"
	"testing"
)

// // import (
// // 	"testing"
// // )

// func TestNewClient(t *testing.T) {
// 	t.Run("connect succeed", func(t *testing.T) {
// 		_, err := New(context.Background())
// 		if err != nil {
// 			t.Error(err)
// 		}
// 	})
// }
// func TesCredentialFileWrite
func TestNewClient(t *testing.T) {
	t.Run("connect succeded", func(t *testing.T) {
		_, err := NewClient(context.Background(), ssm.New())
		if err != nil {
			t.Error(err)
		}
	})
}

// func TestRun(t *testing.T) {
// 	t.Run("display enabled", func(t *testing.T) {
// 		c, _ := NewClient(context.Background(), ssm.New())

// 		service := NewService{}
// 		meta, err := p.DisplayByMetadata(context.Background(), ad.GetInput{
// 			SpaceMetadata: "Qqyqz2cbONJhsFWU_0bTH2x0b1Hsyl2oD_rMVRyREt8",
// 		})
// 		if err != nil {
// 			t.Error(err)
// 		}
// 		if meta == "" {
// 			t.Error("metadata is empty string")
// 		}
// 	})
// }

// type TestService struct {
// }

// func (s *TestService) Upload() {
// }

// func (s *TestService) Download()

// type TestSampleInserter struct {
// }

// func (si *TestSampleInserter) Run(ctx context.Context)(*bigquery.Job, error){
// 	job := *bigquery.Job.
// 	return job, nil
// }

// func TestUpload(t *testing.T){
// 	t.Run("upload succeeded", func(t *testing.T){
// 		service := NewService{}

// 	}
// }
