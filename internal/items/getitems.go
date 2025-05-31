package items

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"strconv"

	"github.com/ehsundar/go-boilerplate/internal/jsonx"
	"github.com/ehsundar/go-boilerplate/internal/storage"
)

type GetItemsResponse struct {
	Items         []storage.Item `json:"items"`
	NextPageToken string         `json:"nextPageToken"`
}

type PageToken struct {
	LastItemID int32 `json:"lastItemId"`
}

func (s *Server) GetItems(w http.ResponseWriter, r *http.Request) {
	pageSize, pageToken, err := getPaginationParams(r)
	if err != nil {
		slog.ErrorContext(r.Context(), "error getting pagination params", "error", err)
		jsonx.WriteError(w, r, http.StatusBadRequest, "invalid pagination params")

		return
	}

	slog.InfoContext(
		r.Context(),
		"pagination params",
		"page_size",
		pageSize,
		"page_token",
		pageToken,
	)

	items, err := s.querier.GetItemsPaginated(r.Context(), s.pool, pageToken.LastItemID, pageSize+1)
	if err != nil {
		slog.ErrorContext(r.Context(), "error getting items", "error", err)
		jsonx.WriteError(w, r, http.StatusInternalServerError, "internal error")

		return
	}

	if items == nil {
		items = make([]storage.Item, 0)
	}

	var responseItems []storage.Item
	if len(items) > int(pageSize) {
		responseItems = items[:pageSize]
	} else {
		responseItems = items
	}

	resp := GetItemsResponse{
		Items:         responseItems,
		NextPageToken: getNextPageToken(items, pageSize),
	}
	jsonx.WriteJSON(w, r, http.StatusOK, resp)
}

// getPaginationParams parses the pagination parameters from the request URL
// based on google aip pagination https://google.aip.dev/158
func getPaginationParams(r *http.Request) (int32, PageToken, error) {
	var pageSize int32

	var pageToken PageToken

	if r.URL.Query().Has("page_size") {
		pageSizeStr := r.URL.Query().Get("page_size")

		pageSizeInt, err := strconv.ParseInt(pageSizeStr, 10, 32)
		if err != nil {
			return 0, PageToken{LastItemID: 0}, fmt.Errorf("parse page_size: %w", err)
		}

		pageSize = int32(pageSizeInt)
	} else {
		pageSize = 10
	}

	if r.URL.Query().Has("page_token") {
		pageTokenStr := r.URL.Query().Get("page_token")

		decoded, err := base64.URLEncoding.DecodeString(pageTokenStr)
		if err != nil {
			return 0, PageToken{LastItemID: 0}, fmt.Errorf("decode page_token: %w", err)
		}

		err = json.Unmarshal(decoded, &pageToken)
		if err != nil {
			return 0, PageToken{LastItemID: 0}, fmt.Errorf("unmarshal page_token: %w", err)
		}
	} else {
		pageToken = PageToken{LastItemID: math.MaxInt32}
	}

	return pageSize, pageToken, nil
}

func getNextPageToken(items []storage.Item, pageSize int32) string {
	if len(items) == 0 {
		return ""
	}

	if len(items) <= int(pageSize) {
		return ""
	}

	lastItem := items[len(items)-2]
	pageToken := PageToken{LastItemID: lastItem.ID}

	encoded, err := json.Marshal(pageToken)
	if err != nil {
		return ""
	}

	return base64.URLEncoding.EncodeToString(encoded)
}
