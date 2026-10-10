package bucket

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteBucketName string

var DeleteBucket = &cobra.Command{
	Use:   "delete",
	Short: "Delete a bucket",
	Long:  "Delete a bucket from the S3 storage.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := s3Client.DeleteBucket(cmd.Context(), deleteBucketName); err != nil {
			return err
		}
		if _, err := fmt.Fprintf(cmd.OutOrStdout(), "Bucket deleted: %s\n", deleteBucketName); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	DeleteBucket.Flags().StringVar(&deleteBucketName, "bucket", "", "Name of the bucket to delete")

	_ = DeleteBucket.MarkFlagRequired("bucket")
}
