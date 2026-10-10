package s3

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (c *client) GetObject(ctx context.Context, bucket string, key string) (io.ReadCloser, error) {
	if err := validateBucketName(bucket); err != nil {
		return nil, err
	}
	result, err := c.s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		return nil, mapError(err)
	}
	return result.Body, nil
}
