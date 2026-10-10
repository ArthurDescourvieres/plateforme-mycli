package object

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/ArthurDescourvieres/plateforme-mycli/cli/internal/s3"
)

type fakeBody struct {
	io.Reader
	closed bool
}

func (b *fakeBody) Close() error {
	b.closed = true
	return nil
}

type fakeS3Client struct {
	putCalled bool
	putBucket string
	putKey    string
	putData   []byte
	getBody   *fakeBody
	getBucket string
	getKey    string
}

func (f *fakeS3Client) CreateBucket(context.Context, string) error {
	return nil
}

func (f *fakeS3Client) DeleteBucket(context.Context, string) error {
	return nil
}

func (f *fakeS3Client) ListBuckets(context.Context) ([]string, error) {
	return nil, nil
}

func (f *fakeS3Client) PutObject(_ context.Context, bucket string, key string, body io.Reader) error {
	f.putCalled = true
	f.putBucket = bucket
	f.putKey = key
	data, err := io.ReadAll(body)
	f.putData = data
	return err
}

func (f *fakeS3Client) GetObject(_ context.Context, bucket string, key string) (io.ReadCloser, error) {
	f.getBucket = bucket
	f.getKey = key
	return f.getBody, nil
}

func (f *fakeS3Client) ListObjects(context.Context, string, string) ([]s3.ObjectItem, error) {
	return nil, nil
}

func (f *fakeS3Client) DeleteObject(context.Context, string, string) error {
	return nil
}

func useFakeClient(t *testing.T, fake *fakeS3Client) {
	t.Helper()
	previousClient := s3Client
	s3Client = fake
	t.Cleanup(func() { s3Client = previousClient })
}

func TestUploadObjectSendsFileContent(t *testing.T) {
	fake := &fakeS3Client{}
	useFakeClient(t, fake)

	content := []byte("upload content")
	path := filepath.Join(t.TempDir(), "report.txt")
	if err := os.WriteFile(path, content, 0o644); err != nil {
		t.Fatalf("write local file: %v", err)
	}

	uploadBucket = "unit-bucket"
	uploadFile = path
	UploadObject.SetOut(io.Discard)

	if err := UploadObject.RunE(UploadObject, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if fake.putBucket != "unit-bucket" {
		t.Errorf("expected bucket unit-bucket, got %q", fake.putBucket)
	}
	if fake.putKey != "report.txt" {
		t.Errorf("expected key report.txt, got %q", fake.putKey)
	}
	if !bytes.Equal(fake.putData, content) {
		t.Errorf("expected content %q, got %q", content, fake.putData)
	}
}

func TestUploadObjectMissingFile(t *testing.T) {
	fake := &fakeS3Client{}
	useFakeClient(t, fake)

	uploadBucket = "unit-bucket"
	uploadFile = filepath.Join(t.TempDir(), "missing.txt")

	err := UploadObject.RunE(UploadObject, nil)
	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("expected a file not found error, got %v", err)
	}
	if fake.putCalled {
		t.Error("PutObject should not be called when the local file is missing")
	}
}

func TestDownloadObjectWritesOutputFile(t *testing.T) {
	content := []byte("download content")
	fake := &fakeS3Client{getBody: &fakeBody{Reader: bytes.NewReader(content)}}
	useFakeClient(t, fake)

	output := filepath.Join(t.TempDir(), "restored.txt")
	downloadBucket = "unit-bucket"
	downloadFile = "report.txt"
	downloadOutput = output
	DownloadObject.SetOut(io.Discard)

	if err := DownloadObject.RunE(DownloadObject, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if fake.getBucket != "unit-bucket" || fake.getKey != "report.txt" {
		t.Errorf("expected unit-bucket/report.txt, got %s/%s", fake.getBucket, fake.getKey)
	}

	got, err := os.ReadFile(output)
	if err != nil {
		t.Fatalf("read output file: %v", err)
	}
	if !bytes.Equal(got, content) {
		t.Errorf("expected content %q, got %q", content, got)
	}
}

func TestDownloadObjectMissingOutputDirectory(t *testing.T) {
	fake := &fakeS3Client{getBody: &fakeBody{Reader: bytes.NewReader(nil)}}
	useFakeClient(t, fake)

	downloadBucket = "unit-bucket"
	downloadFile = "report.txt"
	downloadOutput = filepath.Join(t.TempDir(), "missing", "restored.txt")

	err := DownloadObject.RunE(DownloadObject, nil)
	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("expected a missing directory error, got %v", err)
	}
}

func TestDownloadObjectClosesBody(t *testing.T) {
	body := &fakeBody{Reader: bytes.NewReader([]byte("content"))}
	fake := &fakeS3Client{getBody: body}
	useFakeClient(t, fake)

	downloadBucket = "unit-bucket"
	downloadFile = "report.txt"
	downloadOutput = filepath.Join(t.TempDir(), "restored.txt")
	DownloadObject.SetOut(io.Discard)

	if err := DownloadObject.RunE(DownloadObject, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !body.closed {
		t.Error("the body returned by GetObject should be closed")
	}
}
