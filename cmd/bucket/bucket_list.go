package bucket

import (
	"fmt"

	"github.com/ArthurDescourvieres/plateforme-mycli/internal/s3"
	"github.com/spf13/cobra"
)

var ListBuckets = &cobra.Command{
	Use:   "list",
	Short: "List buckets",
	Run: func(cmd *cobra.Command, args []string) {
		s3.ListBuckets()
		fmt.Fprintln(cmd.OutOrStdout(), "Bucket listing is not implemented yet")
	},
}
