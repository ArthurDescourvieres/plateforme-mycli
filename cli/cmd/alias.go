package cmd

import "github.com/spf13/cobra"

var aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "Manage server connection profiles",
}

func init() {
	rootCmd.AddCommand(aliasCmd)

	aliasCmd.AddCommand(aliasSetCmd)
	aliasCmd.AddCommand(aliasListCmd)
	aliasCmd.AddCommand(aliasUseCmd)
}
