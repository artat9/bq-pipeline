package main

import (
	"context"
	"fmt"
	"kaleido-backend/pkg/infrastructure/s3"
	"os"
)

func main() {
	os.Setenv("EnvironmentId", "v1dev")
	url, err := s3.New().IssueUploadURL(context.Background(), "test/test.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(url)
}
