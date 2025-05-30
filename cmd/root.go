package cmd

import (
	"errors"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var errNoSubcommand = errors.New("please specify a subcommand")

func Execute() error {
	h := slog.NewJSONHandler(os.Stdout, nil)
	slog.SetDefault(slog.New(h))

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
