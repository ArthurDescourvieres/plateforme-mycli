package alias

import "github.com/spf13/cobra"

var (
	aliasSetEndpoint  string
	aliasSetAccessKey string
	aliasSetSecretKey string
)

var SetAlias = &cobra.Command{
	Use:   "set <name>",
	Short: "Set a connection profile",
	Long:  "Create or update a connection profile.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	SetAlias.Flags().StringVar(&aliasSetEndpoint, "endpoint", "", "S3 server endpoint")
	SetAlias.Flags().StringVar(&aliasSetAccessKey, "access-key", "", "S3 access key")
	SetAlias.Flags().StringVar(&aliasSetSecretKey, "secret-key", "", "S3 secret key")

	_ = SetAlias.MarkFlagRequired("endpoint")
	_ = SetAlias.MarkFlagRequired("access-key")
	_ = SetAlias.MarkFlagRequired("secret-key")
}
