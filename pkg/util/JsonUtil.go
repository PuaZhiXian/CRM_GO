package util

import (
	"net/http"
	"encoding/json"
)

func JSON(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if body == nil {
		return
	}

	enc, err := json.Marshal(body)
	if err != nil {
		return
	}
	w.Write(enc)
}
