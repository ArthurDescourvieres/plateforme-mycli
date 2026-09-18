package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var bucketCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a bucket",
	Long:  "Create a new bucket in the S3 storage. This command requires the name of the bucket to be specified using the --bucket flag.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := s3Client.CreateBucket(cmd.Context(), bucketCreateName); err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Bucket created: %s\n", bucketCreateName)
		return nil
	},
}

var bucketCreateName string

func init() {
	bucketCreateCmd.Flags().StringVarP(&bucketCreateName, "bucket", "b", "", "Name of the bucket")
	_ = bucketCreateCmd.MarkFlagRequired("bucket")
	bucketCmd.AddCommand(bucketCreateCmd)
}
