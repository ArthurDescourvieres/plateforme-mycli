package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (c *client) CreateBucket(ctx context.Context, name string) error {
	_, err := c.s3Client.CreateBucket(ctx, &s3.CreateBucketInput{Bucket: &name})
	return err
}

func (c *client) DeleteBucket(ctx context.Context, name string) error {
	_, err := c.s3Client.DeleteBucket(ctx, &s3.DeleteBucketInput{Bucket: &name})
	return err
}

func (c *client) ListBuckets(ctx context.Context) ([]string, error) {
	result, err := c.s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}

	buckets := make([]string, 0, len(result.Buckets))
	for _, bucket := range result.Buckets {
		if bucket.Name != nil {
			buckets = append(buckets, *bucket.Name)
		}
	}
	return buckets, nil
}
