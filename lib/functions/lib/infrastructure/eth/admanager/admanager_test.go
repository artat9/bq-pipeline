package admanager

import (
	"aurora-backend/lib/functions/lib/ad"
	"context"
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
		meta, err := p.DisplayByIndex(context.Background(), ad.GetInput{
			Account: "0xf9B2aAeaFaEE2BfB816B43123962A2fe05Ab1206",
			Index:   1,
		})
		if err != nil {
			t.Error(err)
		}
		if meta == "" {
			t.Error("metadata is empty string")
		}
	})
}
