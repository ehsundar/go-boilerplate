package storage

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
)

//go:embed migrations
var migrationFiles embed.FS

func EnsureMigrationsDone(driver database.Driver, dbName string) error {
	httpFS, err := fs.Sub(migrationFiles, "migrations")
	if err != nil {
		return fmt.Errorf("failed to get subdirectory 'migrations': %w", err)
	}

	srcDriver, err := httpfs.New(http.FS(httpFS), ".")
	if err != nil {
		return fmt.Errorf("failed to create httpfs source driver: %w", err)
	}

	migrateInstance, err := migrate.NewWithInstance(
		"httpfs",
		srcDriver,
		dbName,
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	err = migrateInstance.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	return nil
}
