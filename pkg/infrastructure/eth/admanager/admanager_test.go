package admanager

import (
	"bq-pipeline/pkg/ad"
	"bq-pipeline/pkg/infrastructure/ssm"
	"context"
	"testing"
)

func TestNewProvider(t *testing.T) {
	t.Run("connect succeed", func(t *testing.T) {
		_, err := NewProvider(context.Background(), ssm.New())
		if err != nil {
			t.Error(err)
		}
	})
}

func TestDisplayByIndex(t *testing.T) {
	t.Run("display enabled", func(t *testing.T) {
		p, _ := NewProvider(context.Background(), ssm.New())
		meta, err := p.DisplayByMetadata(context.Background(), ad.GetInput{
			SpaceMetadata: "Qqyqz2cbONJhsFWU_0bTH2x0b1Hsyl2oD_rMVRyREt8",
		})
		if err != nil {
			t.Error(err)
		}
		if meta == "" {
			t.Error("metadata is empty string")
		}
	})
}
