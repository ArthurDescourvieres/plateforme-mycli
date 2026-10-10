package bucket

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/ArthurDescourvieres/plateforme-mycli/cli/internal/config"
	"github.com/ArthurDescourvieres/plateforme-mycli/cli/internal/s3"
)

type fakeS3Client struct {
	createdBucket string
	deletedBucket string
	createError   error
	deleteError   error
}

func (f *fakeS3Client) CreateBucket(_ context.Context, name string) error {
	f.createdBucket = name
	return f.createError
}

func (f *fakeS3Client) DeleteBucket(_ context.Context, name string) error {
	f.deletedBucket = name
	return f.deleteError
}

func (f *fakeS3Client) ListBuckets(context.Context) ([]string, error) {
	return nil, nil
}

func (f *fakeS3Client) PutObject(context.Context, string, string, io.Reader) error {
	return nil
}

func (f *fakeS3Client) GetObject(context.Context, string, string) (io.ReadCloser, error) {
	return nil, nil
}

func (f *fakeS3Client) ListObjects(context.Context, string, string) ([]s3.ObjectItem, error) {
	return nil, nil
}

func (f *fakeS3Client) DeleteObject(context.Context, string, string) error {
	return nil
}

func TestMain(m *testing.M) {
	cfg := config.Load()

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
	defer func() { _ = response.Body.Close() }()

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

	name := fmt.Sprintf("test-bucket-%d", time.Now().UnixNano())
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

	name := fmt.Sprintf("test-bucket-%d", time.Now().UnixNano())
	if err := s3Client.CreateBucket(ctx, name); err != nil {
		t.Fatalf("create bucket for delete test: %v", err)
	}

	if err := s3Client.DeleteBucket(ctx, name); err != nil {
		t.Errorf("delete %q bucket: %v", name, err)
	}
}

func TestCreateBucketCommand(t *testing.T) {
	previousClient := s3Client
	fake := &fakeS3Client{}
	s3Client = fake
	t.Cleanup(func() { s3Client = previousClient })

	if err := CreateBucket.RunE(CreateBucket, []string{"unit-bucket"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if fake.createdBucket != "unit-bucket" {
		t.Errorf("expected unit-bucket, got %q", fake.createdBucket)
	}
}

func TestCreateBucketCommandReturnsError(t *testing.T) {
	previousClient := s3Client
	expectedErr := errors.New("create failed")
	fake := &fakeS3Client{createError: expectedErr}
	s3Client = fake
	t.Cleanup(func() { s3Client = previousClient })

	err := CreateBucket.RunE(CreateBucket, []string{"unit-bucket"})
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected %v, got %v", expectedErr, err)
	}
}

func TestDeleteBucketCommand(t *testing.T) {
	previousClient := s3Client
	fake := &fakeS3Client{}
	s3Client = fake
	t.Cleanup(func() { s3Client = previousClient })

	if err := DeleteBucket.RunE(DeleteBucket, []string{"unit-bucket"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if fake.deletedBucket != "unit-bucket" {
		t.Errorf("expected unit-bucket, got %q", fake.deletedBucket)
	}
}
