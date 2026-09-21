package alias

import "github.com/spf13/cobra"

var ListAliases = &cobra.Command{
	Use:   "list",
	Short: "List connection profiles",
	Long:  "List connection profiles and their endpoints.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
