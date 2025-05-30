package cmd

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"

	"github.com/ehsundar/go-boilerplate/internal/items"
	"github.com/ehsundar/go-boilerplate/internal/storage"
)

const (
	shutdownTimeout   = 10 * time.Second
	readHeaderTimeout = 5 * time.Second
)

func RegisterServeCommand(root *cobra.Command) {
	root.AddCommand(&cobra.Command{
		Use:   "serve",
		Short: "Start the HTTP server",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return serve(cmd.Context())
		},
	})
}

func serve(ctx context.Context) error {
	config, err := LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	pool, err := storage.NewConnectionPool(ctx, config.PostgresConn)
	if err != nil {
		return fmt.Errorf("failed to connect to postgres: %w", err)
	}
	defer pool.Close()

	_, err = storage.NewRedisClient(ctx, config.RedisConn)
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}

	querier := storage.New()
	itemsServer := items.NewServer(pool, querier)
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:              config.ServerAddr,
		Handler:           mux,
		BaseContext:       func(_ net.Listener) context.Context { return ctx },
		ReadHeaderTimeout: readHeaderTimeout,
	}

	registerRoutes(mux, itemsServer)

	go func() {
		slog.Info("Starting server", "address", config.ServerAddr)

		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			slog.Error("HTTP server error", "error", err)
		}

		slog.Info("Stopped serving new connections")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, cancel := context.WithTimeout(ctx, shutdownTimeout)
	defer cancel()

	slog.Info("Shutting down server")

	if err := server.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("failed to shutdown server: %w", err)
	}

	return nil
}

func registerRoutes(mux *http.ServeMux, itemsServer *items.Server) {
	mux.HandleFunc("/items", itemsServer.GetItems)
}
