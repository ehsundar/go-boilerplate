package items

import (
	"log/slog"
	"net/http"

	"github.com/ehsundar/go-boilerplate/internal/jsonx"
)

type CreateItemRequest struct {
	Name string `json:"name"`
}

type CreateItemResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (s *Server) CreateItem(w http.ResponseWriter, r *http.Request) {
	var req CreateItemRequest
	if !jsonx.ReadJSON(w, r, &req) {
		return
	}

	if req.Name == "" {
		jsonx.WriteError(w, r, http.StatusBadRequest, "name is required")

		return
	}

	id, err := s.querier.CreateItem(r.Context(), s.pool, req.Name)
	if err != nil {
		slog.ErrorContext(r.Context(), "error creating item", "error", err)
		jsonx.WriteError(w, r, http.StatusInternalServerError, "internal error")

		return
	}

	resp := CreateItemResponse{ID: id, Name: req.Name}
	jsonx.WriteJSON(w, r, http.StatusCreated, resp)
}
