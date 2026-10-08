package object

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var uploadBucket string
var uploadFile string

var UploadObject = &cobra.Command{
	Use:   "upload",
	Short: "Upload a file into a bucket",
	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := os.Open(uploadFile)
		if err != nil {
			return err
		}
		defer func() { _ = file.Close() }()

		key := filepath.Base(uploadFile)
		err = s3Client.PutObject(cmd.Context(), uploadBucket, key, file)
		if err != nil {
			return err
		}

		if _, err := fmt.Fprintf(cmd.OutOrStdout(), "Object uploaded: %s/%s\n", uploadBucket, key); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	UploadObject.Flags().StringVar(&uploadBucket, "bucket", "", "Name of the target bucket")
	UploadObject.Flags().StringVar(&uploadFile, "file", "", "Path to the local file to upload")

	_ = UploadObject.MarkFlagRequired("bucket")
	_ = UploadObject.MarkFlagRequired("file")
}
