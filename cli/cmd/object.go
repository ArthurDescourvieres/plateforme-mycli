package cmd

import (
	"github.com/ArthurDescourvieres/plateforme-mycli/cli/cmd/object"
	"github.com/spf13/cobra"
)

var objectCmd = &cobra.Command{
	Use:   "object",
	Short: "Manage objects inside a bucket",
}

func init() {
	rootCmd.AddCommand(objectCmd)
	objectCmd.AddCommand(
		object.ListObjects,
		object.UploadObject,
		object.DownloadObject,
		object.DeleteObject,
	)
}
