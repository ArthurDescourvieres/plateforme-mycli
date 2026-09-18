package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var uploadBucket string
var uploadFile string

var objectUploadCmd = &cobra.Command{
	Use: "upload",
	Short: "Upload a file into a bucket",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("not implemented yet")
	},
}

func init() {
	objectUploadCmd.Flags().StringVar(&uploadBucket, "bucket", "", "Name of the target bucket")
	objectUploadCmd.Flags().StringVar(&uploadFile, "file", "", "Path to the local file to upload")

	objectUploadCmd.MarkFlagRequired("bucket")
	objectUploadCmd.MarkFlagRequired("file")

	objectCmd.AddCommand(objectUploadCmd)
}