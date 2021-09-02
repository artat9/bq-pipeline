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
		p.DisplayByIndex(context.Background(), ad.GetInput{})
	})
}
