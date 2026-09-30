package s3

import (
	"errors"
	"fmt"

	"github.com/aws/smithy-go"
)

type S3Error struct {
	Code       string
	HTTPStatus int
	Message    string
}

func (e *S3Error) Error() string {
	if e.Code != "" && e.HTTPStatus != 0 {
		return fmt.Sprintf("[%s] %s (http %d)", e.Code, e.Message, e.HTTPStatus)
	}
	if e.Code != "" {
		return fmt.Sprintf("[%s] %s", e.Code, e.Message)
	}
	return e.Message
}

func mapError(err error) error {
	if err == nil {
		return nil
	}

	var apiErr smithy.APIError
	if errors.As(err, &apiErr) {
		switch apiErr.ErrorCode() {
		case "NoSuchBucket":
			return &S3Error{
				Code:       "NoSuchBucket",
				HTTPStatus: 404,
				Message:    "bucket does not exist; create it first or check the name",
			}
		case "NoSuchKey":
			return &S3Error{
				Code:       "NoSuchKey",
				HTTPStatus: 404,
				Message:    "object does not exist; check the file name in the bucket",
			}
		case "InvalidAccessKeyId":
			return &S3Error{
				Code:       "InvalidAccessKeyId",
				HTTPStatus: 403,
				Message:    "invalid access key; check MYCLI_ACCESS_KEY or your alias",
			}
		case "SignatureDoesNotMatch":
			return &S3Error{
				Code:       "SignatureDoesNotMatch",
				HTTPStatus: 403,
				Message:    "invalid secret key; check MYCLI_SECRET_KEY or your alias",
			}
		case "AccessDenied":
			return &S3Error{
				Code:       "AccessDenied",
				HTTPStatus: 403,
				Message:    "access denied; check your credentials and permissions",
			}
		case "BucketNotEmpty":
			return &S3Error{
				Code:       "BucketNotEmpty",
				HTTPStatus: 409,
				Message:    "bucket is not empty; delete its objects first",
			}
		case "BucketAlreadyExists":
			return &S3Error{
				Code:       "BucketAlreadyExists",
				HTTPStatus: 409,
				Message:    "bucket already exists; choose another name",
			}
		case "BucketAlreadyOwnedByYou":
			return &S3Error{
				Code:       "BucketAlreadyOwnedByYou",
				HTTPStatus: 409,
				Message:    "you already own this bucket; choose another name",
			}
		case "InvalidBucketName":
			return &S3Error{
				Code:       "InvalidBucketName",
				HTTPStatus: 400,
				Message:    "invalid bucket name; use 3-63 lowercase letters, numbers, and hyphens",
			}
		case "XMinioInvalidResourceName":
			return &S3Error{
				Code:       "XMinioInvalidResourceName",
				HTTPStatus: 400,
				Message:    "invalid bucket name; avoid '.' or '..' components",
			}
		default:
			return &S3Error{
				Code:    apiErr.ErrorCode(),
				Message: apiErr.ErrorMessage(),
			}
		}
	}

	return &S3Error{
		Message: "cannot reach the S3 server; is MinIO running?",
	}
}
