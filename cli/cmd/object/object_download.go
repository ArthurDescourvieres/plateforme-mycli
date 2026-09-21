package object

import (
	"fmt"

	"github.com/spf13/cobra"
)

var downloadBucket string
var downloadFile string
var downloadOutput string

var DownloadObject = &cobra.Command{
	Use:   "download",
	Short: "Download an object from a bucket",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("download is not implemented yet")
	},
}

func init() {
	DownloadObject.Flags().StringVar(&downloadBucket, "bucket", "", "Name of the source bucket")
	DownloadObject.Flags().StringVar(&downloadFile, "file", "", "Name of the object to download")
	DownloadObject.Flags().StringVar(&downloadOutput, "output", "", "Local path where the file is written")

	DownloadObject.MarkFlagRequired("bucket")
	DownloadObject.MarkFlagRequired("file")
	DownloadObject.MarkFlagRequired("output")
}
