package cmd

import "github.com/spf13/cobra"

var aliasUseCmd = &cobra.Command{
	Use:   "use <name>",
	Short: "Use a connection profile",
	Long:  "Select a connection profile.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
