package cmd

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4/database/pgx/v5"
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

	postgres := pgx.Postgres{}

	driver, err := postgres.Open(config.PostgresConn)
	if err != nil {
		return fmt.Errorf("failed to open postgres driver: %w", err)
	}

	err = storage.EnsureMigrationsDone(driver, "boilerplate")

	closeErr := driver.Close()

	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	if closeErr != nil {
		return fmt.Errorf("failed to close driver: %w", closeErr)
	}

	return nil
}
