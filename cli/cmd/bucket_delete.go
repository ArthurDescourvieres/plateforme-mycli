package cmd

import (
	"fmt"

	"github.com/ArthurDescourvieres/plateforme-mycli/cli/internal/s3"
	"github.com/spf13/cobra"
)

var bucketDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a bucket",
	Long:  "Delete a bucket from the S3 storage. This command requires the name of the bucket to be specified using the --bucket flag.",
	Run: func(cmd *cobra.Command, args []string) {
		s3.DeleteBucket(bucketDeleteName)
		fmt.Fprintf(cmd.OutOrStdout(), "Bucket deleted: %s\n", bucketDeleteName)
	},
}

var bucketDeleteName string

func init() {
	bucketDeleteCmd.Flags().StringVarP(&bucketDeleteName, "bucket", "b", "", "Name of the bucket")
	_ = bucketDeleteCmd.MarkFlagRequired("bucket")
	bucketCmd.AddCommand(bucketDeleteCmd)
}
