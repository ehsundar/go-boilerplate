package jsonx

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func ReadJSON(w http.ResponseWriter, r *http.Request, v any) bool {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		slog.ErrorContext(r.Context(), "error decoding JSON", "error", err)
		WriteError(w, r, http.StatusBadRequest, "invalid JSON")

		return false
	}

	return true
}

func WriteJSON(w http.ResponseWriter, r *http.Request, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		slog.ErrorContext(r.Context(), "error encoding JSON", "error", err)
		WriteError(w, r, http.StatusInternalServerError, "internal error")
	}
}

func WriteError(w http.ResponseWriter, r *http.Request, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(errorResponse{Error: msg})
	if err != nil {
		slog.ErrorContext(r.Context(), "error writing JSON", "error", err)
	}
}
