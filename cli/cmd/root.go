package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version string = "dev"

var rootCmd = &cobra.Command{
	Use:           "mycli",
	Short:         "Command-line client for an S3 server",
	Version:       version,
	SilenceErrors: true,
	SilenceUsage:  true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
