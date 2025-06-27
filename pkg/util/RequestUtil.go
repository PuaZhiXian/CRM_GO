package util

import (
	"encoding/json"
	"net/http"
)

func DecodeDataFromReq(w http.ResponseWriter, r *http.Request, dataSchema any) bool {
	if err := json.NewDecoder(r.Body).Decode(dataSchema); err != nil {
		RequestErrWriter(w, ErrBadRequest)
		return false
	}
	return true
}
