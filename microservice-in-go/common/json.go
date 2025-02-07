package common

import (
	"encoding/json"
	"net/http"
)

// Marshal JSON
func WriteJSON(rw http.ResponseWriter, status int, data any) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	json.NewEncoder(rw).Encode(data)
}

func ReadJSON(r *http.Request, data any) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func WriteError(rw http.ResponseWriter, status int, message string) {
	WriteJSON(rw, status, map[string]string{"error": message})
}