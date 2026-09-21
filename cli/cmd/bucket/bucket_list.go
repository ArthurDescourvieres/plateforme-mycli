package bucket

import (
	"fmt"

	"github.com/ArthurDescourvieres/plateforme-mycli/cli/internal/s3"
	"github.com/spf13/cobra"
)

var s3Client s3.S3Client

func SetClient(client s3.S3Client) {
	s3Client = client
}

var ListBuckets = &cobra.Command{
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
