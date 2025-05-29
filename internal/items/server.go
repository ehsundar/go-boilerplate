package items

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ehsundar/go-boilerplate/internal/storage"
)

type Server struct {
	pool    *pgxpool.Pool
	querier storage.Querier
}

func NewServer(pool *pgxpool.Pool, querier storage.Querier) *Server {
	return &Server{
		pool:    pool,
		querier: querier,
	}
}
