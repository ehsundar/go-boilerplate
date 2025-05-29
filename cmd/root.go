package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var errNoSubcommand = errors.New("please specify a subcommand")

func Execute() error {
	rootCmd := &cobra.Command{
		Use:   "boilerplate",
		Short: "Boilerplace is a simple REST API application",
		Long:  "Boilerplate is a simple REST API application for demonstration and quick project bootstrapping.",
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
