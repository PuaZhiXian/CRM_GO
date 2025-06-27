package util

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

var ErrBadRequest = errors.New("bad request")
var ErrUnexpected =  errors.New("unexpected error")

type Error struct {
	Code    int
	Message string
}

func writeError(w http.ResponseWriter, msg string, code int) {
	resp := Error{
		Code:    code,
		Message: msg,
	}

	log.Println("Error Msg " + msg)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrWriter = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrHandler = func(w http.ResponseWriter) {
		writeError(w, "Unexpected Err", http.StatusInternalServerError)
	}
)
