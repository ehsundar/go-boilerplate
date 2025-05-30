package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var errNoSubcommand = errors.New("please specify a subcommand")

func Execute() error {
	rootCmd := &cobra.Command{
		Use:   AppName,
		Short: AppName + " is a REST API application",
		RunE: func(_ *cobra.Command, _ []string) error {
			return errNoSubcommand
		},
	}

	RegisterVersionCommand(rootCmd)
	RegisterServeCommand(rootCmd)
	RegisterMigrateCommand(rootCmd)

	//nolint:wrapcheck
	return rootCmd.Execute()
}
