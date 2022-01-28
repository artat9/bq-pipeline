package main

import (
	"bq-pipeline/pkg/infrastructure/s3"
	"context"
	"fmt"
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
