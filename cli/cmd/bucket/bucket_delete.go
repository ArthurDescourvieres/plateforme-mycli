package bucket

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DeleteBucket = &cobra.Command{
	Use:   "delete <bucket>",
	Short: "Delete a bucket",
	Long:  "Delete a bucket from the S3 storage.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := s3Client.DeleteBucket(cmd.Context(), args[0]); err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Bucket deleted: %s\n", args[0])
		return nil
	},
}
