package cmd

import (
	"fmt"

	"github.com/ArthurDescourvieres/plateforme-mycli/cli/internal/s3"
	"github.com/spf13/cobra"
)

var bucketCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a bucket",
	Long:  "Create a new bucket in the S3 storage. This command requires the name of the bucket to be specified using the --bucket flag.",
	Run: func(cmd *cobra.Command, args []string) {
		s3.CreateBucket(bucketCreateName)
		fmt.Fprintf(cmd.OutOrStdout(), "Bucket created: %s\n", bucketCreateName)
	},
}

var bucketCreateName string

func init() {
	bucketCreateCmd.Flags().StringVarP(&bucketCreateName, "bucket", "b", "", "Name of the bucket")
	_ = bucketCreateCmd.MarkFlagRequired("bucket")
	bucketCmd.AddCommand(bucketCreateCmd)
}
