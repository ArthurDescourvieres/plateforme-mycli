package s3

import (
	"bytes"
	"context"
	"errors"
	"strings"
	"testing"
)

func TestValidateBucketNameAcceptsValidNames(t *testing.T) {
	for _, name := range []string{
		"abc",
		"my-bucket",
		"backup.2026",
		"a1-b2.c3",
		strings.Repeat("a", 63),
	} {
		if err := validateBucketName(name); err != nil {
			t.Errorf("validateBucketName(%q): unexpected error: %v", name, err)
		}
	}
}

func TestValidateBucketNameRejectsInvalidNames(t *testing.T) {
	for _, name := range []string{
		"",
		"ab",
		strings.Repeat("a", 64),
		"My-Bucket",
		"my_bucket",
		"my bucket",
		"-bucket",
		"bucket-",
		".bucket",
		"my..bucket",
		"192.168.5.4",
	} {
		err := validateBucketName(name)

		var s3Err *S3Error
		if !errors.As(err, &s3Err) || s3Err.Code != "InvalidBucketName" {
			t.Errorf("validateBucketName(%q): expected an InvalidBucketName error, got %v", name, err)
		}
	}
}

func TestClientRejectsInvalidBucketNameBeforeSendingRequest(t *testing.T) {
	c := &client{}
	ctx := context.Background()
	name := "Invalid_Bucket"

	_, getErr := c.GetObject(ctx, name, "key")
	_, listErr := c.ListObjects(ctx, name, "")

	for operation, err := range map[string]error{
		"CreateBucket": c.CreateBucket(ctx, name),
		"DeleteBucket": c.DeleteBucket(ctx, name),
		"PutObject":    c.PutObject(ctx, name, "key", bytes.NewReader(nil)),
		"GetObject":    getErr,
		"ListObjects":  listErr,
		"DeleteObject": c.DeleteObject(ctx, name, "key"),
	} {
		var s3Err *S3Error
		if !errors.As(err, &s3Err) || s3Err.Code != "InvalidBucketName" {
			t.Errorf("%s: expected an InvalidBucketName error, got %v", operation, err)
		}
	}
}
