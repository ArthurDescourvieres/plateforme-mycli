package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ArthurDescourvieres/plateforme-mycli/cli/cmd"
	"github.com/ArthurDescourvieres/plateforme-mycli/cli/internal/config"
	"github.com/ArthurDescourvieres/plateforme-mycli/cli/internal/s3"
)

func main() {
	client, err := s3.NewClient(context.Background(), config.Load())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cmd.Execute(client)
}
