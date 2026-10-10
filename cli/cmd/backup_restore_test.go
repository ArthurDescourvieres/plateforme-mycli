package cmd

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/ArthurDescourvieres/plateforme-mycli/cli/cmd/bucket"
	"github.com/ArthurDescourvieres/plateforme-mycli/cli/cmd/object"
	"github.com/ArthurDescourvieres/plateforme-mycli/cli/internal/config"
	"github.com/ArthurDescourvieres/plateforme-mycli/cli/internal/s3"
)

func setupClient(t *testing.T) {
	cfg := config.Load()

	if cfg.AccessKey == "" || cfg.SecretKey == "" {
		t.Skip("no S3 credentials: set MYCLI_ACCESS_KEY and MYCLI_SECRET_KEY, or load docker/.env")
	}

	client, err := s3.NewClient(context.Background(), cfg)
	if err != nil {
		t.Fatalf("create S3 client: %v", err)
	}

	bucket.SetClient(client)
	object.SetClient(client)
}

func runCommand(t *testing.T, args ...string) string {
	var output bytes.Buffer
	rootCmd.SetOut(&output)
	rootCmd.SetErr(&output)
	rootCmd.SetArgs(args)

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("mycli %s: %v\n%s", strings.Join(args, " "), err, output.String())
	}

	return output.String()
}

func TestBackupRestore(t *testing.T) {
	setupClient(t)

	bucketName := fmt.Sprintf("backup-test-%d", time.Now().UnixNano())
	dir := t.TempDir()
	original := filepath.Join(dir, "original.txt")
	restored := filepath.Join(dir, "restored.txt")
	content := []byte("MyCLI backup and restore test\n")

	err := os.WriteFile(original, content, 0644)
	if err != nil {
		t.Fatalf("write local file: %v", err)
	}

	runCommand(t, "bucket", "create", "--bucket", bucketName)
	t.Cleanup(func() {
		runCommand(t, "bucket", "delete", "--bucket", bucketName)
	})
	runCommand(t, "object", "upload", "--bucket", bucketName, "--file", original)
	t.Cleanup(func() {
		runCommand(t, "object", "delete", "--bucket", bucketName, "--file", "original.txt")
	})

	listing := runCommand(t, "object", "list", "--bucket", bucketName)
	if !strings.Contains(listing, "original.txt") {
		t.Fatalf("object list does not show original.txt:\n%s", listing)
	}

	runCommand(t, "object", "download", "--bucket", bucketName, "--file", "original.txt", "--output", restored)

	got, err := os.ReadFile(restored)
	if err != nil {
		t.Fatalf("read restored file: %v", err)
	}
	if !bytes.Equal(content, got) {
		t.Fatalf("restored file differs from original:\nwant %q\ngot  %q", content, got)
	}
}
