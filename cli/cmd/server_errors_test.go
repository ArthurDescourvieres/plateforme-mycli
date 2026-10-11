package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/ArthurDescourvieres/plateforme-mycli/cli/internal/s3"
)

func runCommandExpectingError(t *testing.T, args ...string) *s3.S3Error {
	t.Helper()

	var output bytes.Buffer
	rootCmd.SetOut(&output)
	rootCmd.SetErr(&output)
	rootCmd.SetArgs(args)

	err := rootCmd.Execute()

	var s3Err *s3.S3Error
	if !errors.As(err, &s3Err) {
		t.Fatalf("mycli %s: expected an S3 error, got %v\n%s", strings.Join(args, " "), err, output.String())
	}
	return s3Err
}

func TestServerErrors(t *testing.T) {
	setupClient(t)

	missingBucket := fmt.Sprintf("missing-bucket-%d", time.Now().UnixNano())
	if err := runCommandExpectingError(t, "object", "list", "--bucket", missingBucket); err.Code != "NoSuchBucket" {
		t.Errorf("object list on a missing bucket: expected NoSuchBucket, got %s", err.Code)
	}

	bucketName := fmt.Sprintf("errors-test-%d", time.Now().UnixNano())
	runCommand(t, "bucket", "create", "--bucket", bucketName)
	t.Cleanup(func() {
		runCommand(t, "bucket", "delete", "--bucket", bucketName)
	})

	output := filepath.Join(t.TempDir(), "absent.txt")
	if err := runCommandExpectingError(t, "object", "download", "--bucket", bucketName, "--file", "absent.txt", "--output", output); err.Code != "NoSuchKey" {
		t.Errorf("object download of a missing file: expected NoSuchKey, got %s", err.Code)
	}
}
