package bucket

import (
	"fmt"

	"github.com/ArthurDescourvieres/plateforme-mycli/internal/s3"
	"github.com/spf13/cobra"
)

var CreateBucket = &cobra.Command{
	Use:   "create",
	Short: "Create a bucket",
	Long:  "Create a new bucket in the S3 storage. This command requires the name of the bucket to be specified using the --bucket flag.",
	Run: func(cmd *cobra.Command, args []string) {
		s3.CreateBucket(bucketName)
		fmt.Fprintf(cmd.OutOrStdout(), "Bucket created: %s\n", bucketName)
	},
}

var bucketName string

func init() {
	CreateBucket.Flags().StringVarP(&bucketName, "bucket", "b", "", "Name of the bucket")
	_ = CreateBucket.MarkFlagRequired("bucket")
}
