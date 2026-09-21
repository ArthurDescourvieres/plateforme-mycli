package object

import (
	"fmt"

	"github.com/spf13/cobra"
)

var uploadBucket string
var uploadFile string

var UploadObject = &cobra.Command{
	Use: "upload",
	Short: "Upload a file into a bucket",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("not implemented yet")
	},
}

func init() {
	UploadObject.Flags().StringVar(&uploadBucket, "bucket", "", "Name of the target bucket")
	UploadObject.Flags().StringVar(&uploadFile, "file", "", "Path to the local file to upload")

	UploadObject.MarkFlagRequired("bucket")
	UploadObject.MarkFlagRequired("file")
}
