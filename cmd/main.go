package main

import (
	"os"

	"github.com/spf13/cobra"

	"backup-tool/cmd/app/pull"
	"backup-tool/internal/pkg/lib/cmd"
)

var (
	source      string
	destination string
)

func main() {
	if err := start(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

func start() error {
	rootCmd, err := cmd.New(&cmd.Command{
		Use:   "backup",
		Short: "Backup Data with notifiaction mechanism",
		Long:  `Backup Tool application is able to backup traget data and trigger notifiaction`,
	})
	if err != nil {
		return err
	}

	if err := rootCmd.AddCommand(
		&cmd.Command{
			Use:     "pull [source] [destination]",
			Short:   "Pull data from AWS s3 bucket",
			Long:    `Pull data from AWS s3 bucket`,
			Args:    cobra.MinimumNArgs(2),
			Service: pull.New(),
		},
		&cmd.Command{
			Use:   "push",
			Short: "Push data to AWS s3 bucket",
			Long:  `Push data to AWS s3 bucket`,
		},
	); err != nil {
		return err
	}

	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}
