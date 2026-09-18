package bucket

import (
	"fmt"

	"github.com/ArthurDescourvieres/plateforme-mycli/cli/internal/s3"
	"github.com/spf13/cobra"
)

var DeleteBucket = &cobra.Command{
	Use:   "delete",
	Short: "Delete a bucket",
	Long:  "Delete a bucket from the S3 storage. This command requires the name of the bucket to be specified using the --bucket flag.",
	Run: func(cmd *cobra.Command, args []string) {
		s3.DeleteBucket(bucketName)
		fmt.Fprintf(cmd.OutOrStdout(), "Bucket deleted: %s\n", bucketName)
	},
}

func init() {
	DeleteBucket.Flags().StringVarP(&bucketName, "bucket", "b", "", "Name of the bucket")
	_ = DeleteBucket.MarkFlagRequired("bucket")
}
