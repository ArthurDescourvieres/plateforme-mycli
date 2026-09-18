package bucket

import (
	"fmt"

	"github.com/ArthurDescourvieres/plateforme-mycli/internal/s3"
	"github.com/spf13/cobra"
)

var DeleteBucket = &cobra.Command{
	Use:   "delete",
	Short: "Delete a bucket",
	Run: func(cmd *cobra.Command, args []string) {
		s3.DeleteBucket(bucketName)
		fmt.Fprintf(cmd.OutOrStdout(), "Bucket deletion is not implemented yet: %s\n", bucketName)
	},
}

func init() {
	DeleteBucket.Flags().StringVarP(&bucketName, "bucket", "b", "", "Name of the bucket")
	_ = DeleteBucket.MarkFlagRequired("bucket")
}
