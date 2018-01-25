package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func unmarshal(r io.ReadCloser, object interface{}) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&object)
	defer r.Close()

	return err
}

func response(w http.ResponseWriter, code int, object ...interface{}) {
	if len(object) != 0 {
		if err := json.NewEncoder(w).Encode(object); err != nil {
			w.WriteHeader(500)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
}
