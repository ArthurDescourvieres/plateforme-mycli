package bucket

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/ArthurDescourvieres/plateforme-mycli/cli/internal/config"
	"github.com/ArthurDescourvieres/plateforme-mycli/cli/internal/s3"
)

func TestMain(m *testing.M) {
	cfg := config.Load()
	cfg.AccessKey = "admin"
	cfg.SecretKey = "password"

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		cfg.URL+"/minio/health/ready",
		nil,
	)
	if err != nil {
		panic(err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic("MinIO unreachable: " + err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		panic(fmt.Sprintf(
			"MinIO is not ready: HTTP %d",
			response.StatusCode,
		))
	}

	client, err := s3.NewClient(ctx, cfg)
	if err != nil {
		panic("create S3 client: " + err.Error())
	}
	s3Client = client

	os.Exit(m.Run())
}

func TestCreateBucket(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	name := "test-bucket"
	t.Cleanup(func() {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cleanupCancel()

		if err := s3Client.DeleteBucket(cleanupCtx, name); err != nil {
			t.Logf("cleanup bucket %q: %v", name, err)
		}
	})

	if err := s3Client.CreateBucket(ctx, name); err != nil {
		t.Fatalf("create bucket: %v", err)
	}

	buckets, err := s3Client.ListBuckets(ctx)
	if err != nil {
		t.Fatalf("list buckets: %v", err)
	}

	for _, bucket := range buckets {
		if bucket == name {
			return
		}
	}

	t.Errorf("bucket %q was not created", name)
}

func TestDeleteBucket(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	name := "test-bucket"

	if err := s3Client.DeleteBucket(ctx, name); err != nil {
		t.Errorf("Can't delete %q bucket", name)
	}
}
