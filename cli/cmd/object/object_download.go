package object

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var downloadBucket string
var downloadFile string
var downloadOutput string

var DownloadObject = &cobra.Command{
	Use:   "download",
	Short: "Download an object from a bucket",
	RunE: func(cmd *cobra.Command, args []string) error {
		body, err := s3Client.GetObject(cmd.Context(), downloadBucket, downloadFile)
		if err != nil {
			return err
		}
		defer body.Close()

		file, err := os.Create(downloadOutput)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(file, body)
		if err != nil {
			return err
		}

		fmt.Fprintf(cmd.OutOrStdout(), "Object downloaded: %s/%s -> %s\n", downloadBucket, downloadFile, downloadOutput)
		return nil
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
