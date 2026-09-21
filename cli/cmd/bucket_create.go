package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var bucketCreateCmd = &cobra.Command{
	Use:   "create <bucket>",
	Short: "Create a bucket",
	Long:  "Create a new bucket in the S3 storage.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := s3Client.CreateBucket(cmd.Context(), args[0]); err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Bucket created: %s\n", args[0])
		return nil
	},
}

func init() {
	bucketCmd.AddCommand(bucketCreateCmd)
}
