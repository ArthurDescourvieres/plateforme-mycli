package cmd

import "github.com/spf13/cobra"

var (
 objectDeleteBucket string
 objectDeleteFile string
)

var objectDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete an object from a bucket",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	objectDeleteCmd.Flags().StringVar(&objectDeleteBucket,"bucket", "" ,"bucket name")
	objectDeleteCmd.Flags().StringVar(&objectDeleteFile,"file", "", "object key / file name")
	_ = objectDeleteCmd.MarkFlagRequired("bucket")
	_ = objectDeleteCmd.MarkFlagRequired("file")
	objectCmd.AddCommand(objectDeleteCmd)
}