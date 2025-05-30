package storage

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"github.com/jackc/pgx/v5/pgconn"
)

//go:embed migrations
var migrationFiles embed.FS

func EnsureMigrationsDone(postgresConn string) error {
	httpFS, err := fs.Sub(migrationFiles, "migrations")
	if err != nil {
		return fmt.Errorf("failed to get subdirectory 'migrations': %w", err)
	}

	srcDriver, err := httpfs.New(http.FS(httpFS), ".")
	if err != nil {
		return fmt.Errorf("failed to create httpfs source driver: %w", err)
	}

	pgxDriver, err := (&pgx.Postgres{}).Open(postgresConn)
	if err != nil {
		return fmt.Errorf("failed to open postgres driver: %w", err)
	}
	//nolint:errcheck
	defer pgxDriver.Close()

	cfg, err := pgconn.ParseConfig(postgresConn)
	if err != nil {
		return fmt.Errorf("failed to parse postgres connection string: %w", err)
	}

	migrateInstance, err := migrate.NewWithInstance(
		"httpfs",
		srcDriver,
		cfg.Database,
		pgxDriver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	err = migrateInstance.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		slog.Info("No new migrations to apply.")
	} else {
		slog.Info("Migrations applied successfully.")
	}

	return nil
}
