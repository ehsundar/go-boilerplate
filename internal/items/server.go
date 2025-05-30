package items

import (
	"fmt"
	"net/http"

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

func (s *Server) RegisterRoutes(mux *http.ServeMux, prefix string) {
	mux.HandleFunc(
		"GET "+prefix,
		s.GetItems,
	)
	mux.HandleFunc(
		"POST "+prefix,
		s.CreateItem,
	)
	mux.HandleFunc(
		fmt.Sprintf("GET %s/{id}", prefix),
		s.GetItem,
	)
}
