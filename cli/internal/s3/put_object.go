package s3

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (c *client) PutObject(ctx context.Context, bucket string, key string,
	body io.Reader) error {
	if err := validateBucketName(bucket); err != nil {
		return err
	}
	_, err := c.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   body,
	})
	return mapError(err)
}
