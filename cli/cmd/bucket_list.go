package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var bucketListCmd = &cobra.Command{
	Use:   "list",
	Short: "List buckets",
	Long:  "List all buckets in the S3 storage.",
	RunE: func(cmd *cobra.Command, args []string) error {
		buckets, err := s3Client.ListBuckets(cmd.Context())
		if err != nil {
			return err
		}
		for _, bucket := range buckets {
			fmt.Fprintln(cmd.OutOrStdout(), bucket)
		}
		return nil
	},
}

func init() {
	bucketCmd.AddCommand(bucketListCmd)
}
