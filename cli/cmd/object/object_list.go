package object

import "github.com/spf13/cobra"

var objectListBucket string

var ListObjects = &cobra.Command{
	Use:   "list",
	Short: "list objects in a bucket",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	ListObjects.Flags().StringVar(&objectListBucket, "bucket", "", "bucket name")
	_ = ListObjects.MarkFlagRequired("bucket")
}
