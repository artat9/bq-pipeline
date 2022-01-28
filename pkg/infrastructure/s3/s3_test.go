package s3

import (
	"context"
	"os"
	"strings"
	"testing"
)

func TestS3_IssueUploadURL(t *testing.T) {
	os.Setenv("EnvironmentId", "v1dev")
	t.Run("issue presign success", func(t *testing.T) {
		url, err := New().IssueUploadURL(context.Background(), "test.txt")
		if err != nil {
			t.Error(err)
		}
		if !strings.HasPrefix(url, "https://bq-pipeline-asset") {
			t.Error("got:", url)
		}
	})
}
