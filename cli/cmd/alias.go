package cmd

import (
	"github.com/ArthurDescourvieres/plateforme-mycli/cli/cmd/alias"
	"github.com/spf13/cobra"
)

var aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "Manage server connection profiles",
}

func init() {
	rootCmd.AddCommand(aliasCmd)
	aliasCmd.AddCommand(
		alias.SetAlias,
		alias.ListAliases,
		alias.UseAlias,
	)
}
