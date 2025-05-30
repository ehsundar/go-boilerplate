package items

import (
	"database/sql"
	"errors"
	"log/slog"
	"net/http"

	"github.com/ehsundar/go-boilerplate/internal/jsonx"
)

func (s *Server) GetItem(w http.ResponseWriter, r *http.Request) {
	id, ok := jsonx.ParseID(w, r)
	if !ok {
		return
	}

	item, err := s.querier.GetItem(r.Context(), s.pool, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			jsonx.WriteError(w, r, http.StatusNotFound, "item not found")

			return
		}

		slog.ErrorContext(r.Context(), "error getting item", "error", err)
		jsonx.WriteError(w, r, http.StatusInternalServerError, "internal error")

		return
	}

	jsonx.WriteJSON(w, r, http.StatusOK, item)
}
