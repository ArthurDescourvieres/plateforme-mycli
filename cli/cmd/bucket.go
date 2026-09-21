package cmd

import (
	"github.com/ArthurDescourvieres/plateforme-mycli/cli/cmd/bucket"
	"github.com/spf13/cobra"
)

var bucketCmd = &cobra.Command{
	Use:   "bucket",
	Short: "Manage buckets",
}

func init() {
	rootCmd.AddCommand(bucketCmd)
	bucketCmd.AddCommand(
		bucket.CreateBucket,
		bucket.DeleteBucket,
		bucket.ListBuckets,
	)
}
