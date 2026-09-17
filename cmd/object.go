package cmd

import "github.com/spf13/cobra"

var objectCmd = &cobra.Command{
	Use:   "object",
	Short: "Manage objects inside a bucket",
}

func init() {
	rootCmd.AddCommand(objectCmd)
}