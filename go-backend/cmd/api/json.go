package main

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, data any) error {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return nil
}

func readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	// limit the body size to 1 MB
	maxBytes := 1_048_578
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	// disallow unknown fields
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

type JSONError struct {
	Error string `json:"error"`
}

func writeJSONError(w http.ResponseWriter, status int, message string) error {
	env := &JSONError{Error: message}
	return writeJSON(w, status, env)
}
