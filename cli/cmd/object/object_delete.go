package object

import "github.com/spf13/cobra"

var (
 objectDeleteBucket string
 objectDeleteFile string
)

var DeleteObject = &cobra.Command{
	Use:   "delete",
	Short: "delete an object from a bucket",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	DeleteObject.Flags().StringVar(&objectDeleteBucket,"bucket", "" ,"bucket name")
	DeleteObject.Flags().StringVar(&objectDeleteFile,"file", "", "object key / file name")
	_ = DeleteObject.MarkFlagRequired("bucket")
	_ = DeleteObject.MarkFlagRequired("file")
}