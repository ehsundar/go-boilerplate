package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/ehsundar/go-boilerplate/internal/storage"
)

func RegisterMigrateCommand(root *cobra.Command) {
	root.AddCommand(&cobra.Command{
		Use:   "migrate",
		Short: "Migrate the database",
		Long:  "Run all pending database migrations to keep the schema up to date.",
		RunE: func(_ *cobra.Command, _ []string) error {
			return migrate()
		},
	})
}

func migrate() error {
	config, err := LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	err = storage.EnsureMigrationsDone(config.PostgresConn)
	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}
