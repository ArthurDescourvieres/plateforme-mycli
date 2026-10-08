package object

import (
	"fmt"

	"github.com/ArthurDescourvieres/plateforme-mycli/cli/internal/s3"
	"github.com/spf13/cobra"
)

var s3Client s3.S3Client

func SetClient(client s3.S3Client) {
	s3Client = client
}

var objectListBucket string

var ListObjects = &cobra.Command{
	Use:   "list",
	Short: "List objects in a bucket",
	RunE: func(cmd *cobra.Command, args []string) error {
		items, err := s3Client.ListObjects(cmd.Context(), objectListBucket)
		if err != nil {
			return err
		}
		for _, item := range items {
			if _, err := fmt.Fprintln(cmd.OutOrStdout(), item.Key); err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	ListObjects.Flags().StringVar(&objectListBucket, "bucket", "", "bucket name")
	_ = ListObjects.MarkFlagRequired("bucket")
}
