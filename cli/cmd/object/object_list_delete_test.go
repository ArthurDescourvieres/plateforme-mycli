package object

import (
	"bytes"
	"errors"
	"testing"

	"github.com/ArthurDescourvieres/plateforme-mycli/cli/internal/s3"
)

func TestListObjectsPrintsKeys(t *testing.T) {
	fake := &fakeS3Client{listItems: []s3.ObjectItem{{Key: "reports/a.txt"}, {Key: "reports/b.txt"}}}
	useFakeClient(t, fake)

	var output bytes.Buffer
	ListObjects.SetOut(&output)
	objectListBucket = "unit-bucket"
	objectListPrefix = "reports/"

	if err := ListObjects.RunE(ListObjects, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if fake.listBucket != "unit-bucket" || fake.listPrefix != "reports/" {
		t.Errorf("expected unit-bucket and prefix reports/, got %q and %q", fake.listBucket, fake.listPrefix)
	}
	if got, want := output.String(), "reports/a.txt\nreports/b.txt\n"; got != want {
		t.Errorf("expected output %q, got %q", want, got)
	}
}

func TestListObjectsEmptyBucket(t *testing.T) {
	useFakeClient(t, &fakeS3Client{})

	var output bytes.Buffer
	ListObjects.SetOut(&output)
	objectListBucket = "unit-bucket"
	objectListPrefix = ""

	if err := ListObjects.RunE(ListObjects, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if output.Len() != 0 {
		t.Errorf("expected no output, got %q", output.String())
	}
}

func TestListObjectsReturnsClientError(t *testing.T) {
	expectedErr := &s3.S3Error{Code: "NoSuchBucket", HTTPStatus: 404, Message: "bucket does not exist"}
	useFakeClient(t, &fakeS3Client{listErr: expectedErr})

	var output bytes.Buffer
	ListObjects.SetOut(&output)
	objectListBucket = "missing-bucket"
	objectListPrefix = ""

	err := ListObjects.RunE(ListObjects, nil)
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected %v, got %v", expectedErr, err)
	}
	if output.Len() != 0 {
		t.Errorf("expected no output, got %q", output.String())
	}
}

func TestDeleteObjectSendsBucketAndKey(t *testing.T) {
	fake := &fakeS3Client{}
	useFakeClient(t, fake)

	var output bytes.Buffer
	DeleteObject.SetOut(&output)
	objectDeleteBucket = "unit-bucket"
	objectDeleteFile = "report.txt"

	if err := DeleteObject.RunE(DeleteObject, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if fake.deleteBucket != "unit-bucket" || fake.deleteKey != "report.txt" {
		t.Errorf("expected unit-bucket/report.txt, got %s/%s", fake.deleteBucket, fake.deleteKey)
	}
	if got, want := output.String(), "Object deleted: unit-bucket/report.txt\n"; got != want {
		t.Errorf("expected output %q, got %q", want, got)
	}
}

func TestDeleteObjectReturnsClientError(t *testing.T) {
	expectedErr := &s3.S3Error{Code: "NoSuchBucket", HTTPStatus: 404, Message: "bucket does not exist"}
	useFakeClient(t, &fakeS3Client{deleteErr: expectedErr})

	var output bytes.Buffer
	DeleteObject.SetOut(&output)
	objectDeleteBucket = "missing-bucket"
	objectDeleteFile = "report.txt"

	err := DeleteObject.RunE(DeleteObject, nil)
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected %v, got %v", expectedErr, err)
	}
	if output.Len() != 0 {
		t.Errorf("expected no success message, got %q", output.String())
	}
}
