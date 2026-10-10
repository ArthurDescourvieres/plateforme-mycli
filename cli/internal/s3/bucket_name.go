package s3

import (
	"net"
	"regexp"
	"strings"
)

var bucketNamePattern = regexp.MustCompile(`^[a-z0-9][a-z0-9.-]*[a-z0-9]$`)

func invalidBucketName(message string) error {
	return &S3Error{Code: "InvalidBucketName", Message: message}
}

func validateBucketName(name string) error {
	if len(name) < 3 || len(name) > 63 {
		return invalidBucketName("bucket name must be between 3 and 63 characters long")
	}
	if !bucketNamePattern.MatchString(name) {
		return invalidBucketName("bucket name must use only lowercase letters, numbers, dots and hyphens, and start and end with a letter or a number")
	}
	if strings.Contains(name, "..") {
		return invalidBucketName("bucket name must not contain two adjacent dots")
	}
	if net.ParseIP(name) != nil {
		return invalidBucketName("bucket name must not be formatted as an IP address")
	}
	return nil
}
