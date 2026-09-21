package cmd

import (
	"fmt"
	"os"

	"github.com/ArthurDescourvieres/plateforme-mycli/cli/cmd/bucket"
	"github.com/ArthurDescourvieres/plateforme-mycli/cli/internal/s3"
	"github.com/spf13/cobra"
)

var s3Client s3.S3Client

var rootCmd = &cobra.Command{
	Use:   "mycli",
	Short: "Command-line client for an S3 server",
}

func Execute(client s3.S3Client) {
	s3Client = client
	bucket.SetClient(client)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
