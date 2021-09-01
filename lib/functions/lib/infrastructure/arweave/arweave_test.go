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

func TestDownloadMetadata(t *testing.T) {
	t.Run("enable to download", func(t *testing.T) {
		client, err := NewClient()
		if err != nil {
			t.Error(err)
		}
		res, err := client.downloadMetadata(context.Background(), "_ojROEcQk1EJTEwYAovVBT0uh6t-X7YTDzHjxYLvCfY")
		if err != nil {
			t.Error(err)
		}
		t.Error(res.Description)
	})
}
func TestPost(t *testing.T) {
	client, _ := NewClient()
	t.Run("enable to fetch post content", func(t *testing.T) {
		post, err := client.Post(context.Background(), "ykgGaLOubSDMIe0GYjkkTRruRsGGmaTicDmoyggxq9o")
		if err != nil {
			t.Error(err)
		}
		if post.ProjectName != "Test-2021-08-24" {
			t.Error("pjname not loaded. got ", post.ProjectName)
		}
		if len(post.Image) == 0 {
			t.Error("image not loaded")
		}
		if post.Description == "" {
			t.Error("description not loaded")
		}
	})
}
