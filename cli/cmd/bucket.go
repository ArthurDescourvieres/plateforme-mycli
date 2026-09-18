package cmd

import "github.com/spf13/cobra"

var bucketCmd = &cobra.Command{
	Use:   "bucket",
	Short: "Manage buckets",
}

func init() {
	rootCmd.AddCommand(bucketCmd)
}
