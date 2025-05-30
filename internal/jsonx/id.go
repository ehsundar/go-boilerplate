package jsonx

import (
	"math"
	"net/http"
	"strconv"
)

func ParseID(w http.ResponseWriter, r *http.Request) (int32, bool) {
	value := r.PathValue("id")

	id, err := strconv.ParseInt(value, 10, 32)
	if err != nil || id < 1 || id > math.MaxInt32 {
		WriteError(w, r, http.StatusBadRequest, "invalid id")

		return 0, false
	}

	return int32(id), true
}
