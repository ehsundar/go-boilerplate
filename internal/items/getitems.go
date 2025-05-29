package items

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func (s *Server) GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := s.querier.GetItems(r.Context(), s.pool)
	if err != nil {
		slog.ErrorContext(r.Context(), "error getting items", "error", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		slog.ErrorContext(r.Context(), "error encoding items", "error", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}
