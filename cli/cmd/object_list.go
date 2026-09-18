package cmd

import "github.com/spf13/cobra"

var objectListBucket string

var objectListCmd = &cobra.Command{
	Use:   "list",
	Short: "list objects in a bucket",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	objectListCmd.Flags().StringVar(&objectListBucket, "bucket", "", "bucket name")
	_ = objectListCmd.MarkFlagRequired("bucket")
	objectCmd.AddCommand(objectListCmd)
}
