package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (c *client) DeleteObject(ctx context.Context, bucket, key string) error {
	_, err := c.s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	return mapError(err)
}
