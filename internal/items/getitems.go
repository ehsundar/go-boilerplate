package items

import (
	"log/slog"
	"net/http"

	"github.com/ehsundar/go-boilerplate/internal/jsonx"
	"github.com/ehsundar/go-boilerplate/internal/storage"
)

type ItemsResponse struct {
	Items []storage.Item `json:"items"`
}

func (s *Server) GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := s.querier.GetItems(r.Context(), s.pool)
	if err != nil {
		slog.ErrorContext(r.Context(), "error getting items", "error", err)
		jsonx.WriteError(w, r, http.StatusInternalServerError, "internal error")

		return
	}

	if items == nil {
		items = make([]storage.Item, 0)
	}

	resp := ItemsResponse{Items: items}
	jsonx.WriteJSON(w, r, http.StatusOK, resp)
}
