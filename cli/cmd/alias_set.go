package cmd

import "github.com/spf13/cobra"

var (
	aliasSetEndpoint  string
	aliasSetAccessKey string
	aliasSetSecretKey string
)

var aliasSetCmd = &cobra.Command{
	Use:   "set <name>",
	Short: "Set a connection profile",
	Long:  "Create or update a connection profile.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	aliasSetCmd.Flags().StringVar(&aliasSetEndpoint, "endpoint", "", "S3 server endpoint")
	aliasSetCmd.Flags().StringVar(&aliasSetAccessKey, "access-key", "", "S3 access key")
	aliasSetCmd.Flags().StringVar(&aliasSetSecretKey, "secret-key", "", "S3 secret key")

	_ = aliasSetCmd.MarkFlagRequired("endpoint")
	_ = aliasSetCmd.MarkFlagRequired("access-key")
	_ = aliasSetCmd.MarkFlagRequired("secret-key")
}
