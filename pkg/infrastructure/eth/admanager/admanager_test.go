package admanager

import (
	"context"
	"kaleido-backend/pkg/ad"
	"testing"
)

func TestNewProvider(t *testing.T) {
	t.Run("connect succeed", func(t *testing.T) {
		_, err := NewProvider()
		if err != nil {
			t.Error(err)
		}
	})
}

func TestDisplayByIndex(t *testing.T) {
	t.Run("display enabled", func(t *testing.T) {
		p, _ := NewProvider()
		meta, err := p.DisplayByMetadata(context.Background(), ad.GetInput{
			Account:  "0xb5bE22F33D8f0b1Cc131674C562069D1B5912147",
			Metadata: "Qqyqz2cbONJhsFWU_0bTH2x0b1Hsyl2oD_rMVRyREt8",
		})
		if err != nil {
			t.Error(err)
		}
		if meta == "" {
			t.Error("metadata is empty string")
		}
	})
}
