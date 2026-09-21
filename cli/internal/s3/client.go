package s3

import (
	"context"

	"github.com/ArthurDescourvieres/plateforme-mycli/cli/internal/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client interface {
	CreateBucket(context.Context, string) error
	DeleteBucket(context.Context, string) error
	ListBuckets(context.Context) ([]string, error)
}

type client struct {
	s3Client *s3.Client
}

func NewClient(ctx context.Context, cfg config.Config) (S3Client, error) {
	awsCfg, err := awsConfig.LoadDefaultConfig(
		ctx,
		awsConfig.WithRegion(cfg.Region),
		awsConfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(cfg.AccessKey, cfg.SecretKey, ""),
		),
	)
	if err != nil {
		return nil, err
	}

	return &client{
		s3Client: s3.NewFromConfig(awsCfg, func(options *s3.Options) {
			options.BaseEndpoint = aws.String(cfg.URL)
			options.UsePathStyle = true
		}),
	}, nil
}
