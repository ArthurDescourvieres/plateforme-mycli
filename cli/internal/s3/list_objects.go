package s3

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type ObjectItem struct {
	Key          string
	Size         int64
	LastModified string
}

func (c *client) ListObjects(ctx context.Context, bucket string) ([]ObjectItem, error) {
	result, err := c.s3Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: &bucket,
	})
	if err != nil {
		return nil, mapError(err)
	}

	items := make([]ObjectItem, 0, len(result.Contents))
	for _, obj := range result.Contents {
		item := ObjectItem{}
		if obj.Key != nil {
			item.Key = *obj.Key
		}
		if obj.Size != nil {
			item.Size = *obj.Size
		}
		if obj.LastModified != nil {
			item.LastModified = obj.LastModified.UTC().Format(time.RFC3339)
		}
		items = append(items, item)
	}
	return items, nil
}
