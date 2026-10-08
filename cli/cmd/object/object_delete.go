package object

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	objectDeleteBucket string
	objectDeleteFile   string
)

var DeleteObject = &cobra.Command{
	Use:   "delete",
	Short: "Delete an object from a bucket",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := s3Client.DeleteObject(cmd.Context(), objectDeleteBucket, objectDeleteFile); err != nil {
			return err
		}
		if _, err := fmt.Fprintf(cmd.OutOrStdout(), "Object deleted: %s/%s\n", objectDeleteBucket, objectDeleteFile); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	DeleteObject.Flags().StringVar(&objectDeleteBucket, "bucket", "", "bucket name")
	DeleteObject.Flags().StringVar(&objectDeleteFile, "file", "", "object key / file name")
	_ = DeleteObject.MarkFlagRequired("bucket")
	_ = DeleteObject.MarkFlagRequired("file")
}
