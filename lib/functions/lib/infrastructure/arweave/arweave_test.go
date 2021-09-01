package arweave

import (
	"context"
	"testing"
)

func TestDownload(t *testing.T) {
	t.Run("enable to download", func(t *testing.T) {
		client, err := NewClient()
		if err != nil {
			t.Error(err)
		}
		loaded, err := client.download(context.Background(), "ykgGaLOubSDMIe0GYjkkTRruRsGGmaTicDmoyggxq9o")
		if err != nil {
			t.Error(err)
		}
		if loaded == "" {
			t.Error("content not donlowaded")
		}
	})
}
