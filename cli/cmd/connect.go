package cmd

import (
	"github.com/ArthurDescourvieres/plateforme-mycli/cli/cmd/bucket"
	"github.com/ArthurDescourvieres/plateforme-mycli/cli/cmd/object"
	"github.com/ArthurDescourvieres/plateforme-mycli/cli/internal/config"
	"github.com/ArthurDescourvieres/plateforme-mycli/cli/internal/s3"
	"github.com/spf13/cobra"
)

func connect(cmd *cobra.Command, args []string) error {
	client, err := s3.NewClient(cmd.Context(), config.Load())
	if err != nil {
		return err
	}
	bucket.SetClient(client)
	object.SetClient(client)
	return nil
}
