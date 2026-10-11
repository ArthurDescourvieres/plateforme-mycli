package s3

import (
	"errors"
	"testing"

	"github.com/aws/smithy-go"
)

func TestMapErrorKnownCodes(t *testing.T) {
	for _, tc := range []struct {
		code   string
		status int
	}{
		{"NoSuchBucket", 404},
		{"NoSuchKey", 404},
		{"InvalidAccessKeyId", 403},
		{"SignatureDoesNotMatch", 403},
		{"AccessDenied", 403},
		{"BucketNotEmpty", 409},
		{"BucketAlreadyExists", 409},
		{"BucketAlreadyOwnedByYou", 409},
		{"InvalidBucketName", 400},
		{"XMinioInvalidResourceName", 400},
	} {
		err := mapError(&smithy.GenericAPIError{Code: tc.code, Message: "server message"})

		var s3Err *S3Error
		if !errors.As(err, &s3Err) {
			t.Errorf("%s: expected an *S3Error, got %T", tc.code, err)
			continue
		}
		if s3Err.Code != tc.code || s3Err.HTTPStatus != tc.status {
			t.Errorf("%s: expected code %s and status %d, got %s and %d", tc.code, tc.code, tc.status, s3Err.Code, s3Err.HTTPStatus)
		}
		if s3Err.Message == "" || s3Err.Message == "server message" {
			t.Errorf("%s: expected a translated message, got %q", tc.code, s3Err.Message)
		}
	}
}

func TestMapErrorNil(t *testing.T) {
	if err := mapError(nil); err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
}

func TestMapErrorKeepsUnknownCode(t *testing.T) {
	err := mapError(&smithy.GenericAPIError{Code: "SlowDown", Message: "Please reduce your request rate."})

	var s3Err *S3Error
	if !errors.As(err, &s3Err) {
		t.Fatalf("expected an *S3Error, got %T", err)
	}
	if s3Err.Code != "SlowDown" || s3Err.Message != "Please reduce your request rate." {
		t.Errorf("expected the server code and message, got %q and %q", s3Err.Code, s3Err.Message)
	}
}

func TestMapErrorWithoutServerResponse(t *testing.T) {
	err := mapError(errors.New("dial tcp 127.0.0.1:9000: connection refused"))

	var s3Err *S3Error
	if !errors.As(err, &s3Err) {
		t.Fatalf("expected an *S3Error, got %T", err)
	}
	if s3Err.Code != "" || s3Err.Message != "cannot reach the S3 server; is MinIO running?" {
		t.Errorf("unexpected error: code %q, message %q", s3Err.Code, s3Err.Message)
	}
}

func TestS3ErrorFormats(t *testing.T) {
	for _, tc := range []struct {
		err  S3Error
		want string
	}{
		{S3Error{Code: "NoSuchBucket", HTTPStatus: 404, Message: "bucket does not exist"}, "[NoSuchBucket] bucket does not exist (http 404)"},
		{S3Error{Code: "InvalidBucketName", Message: "bucket name is invalid"}, "[InvalidBucketName] bucket name is invalid"},
		{S3Error{Message: "cannot reach the S3 server"}, "cannot reach the S3 server"},
	} {
		if got := tc.err.Error(); got != tc.want {
			t.Errorf("expected %q, got %q", tc.want, got)
		}
	}
}
