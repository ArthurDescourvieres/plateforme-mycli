package bucket

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createBucketName string

var CreateBucket = &cobra.Command{
	Use:   "create",
	Short: "Create a bucket",
	Long:  "Create a new bucket in the S3 storage.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := s3Client.CreateBucket(cmd.Context(), createBucketName); err != nil {
			return err
		}
		if _, err := fmt.Fprintf(cmd.OutOrStdout(), "Bucket created: %s\n", createBucketName); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	CreateBucket.Flags().StringVar(&createBucketName, "bucket", "", "Name of the bucket to create")

	_ = CreateBucket.MarkFlagRequired("bucket")
}
