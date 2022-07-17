package utils

import (
	"encoding/json"
	"net/http"
)

// Writes a Json object containing an error message to the response writer, and sets HTTP status code
// This is an alternative to http.error (which writes text only).
func WriteJSONError(w http.ResponseWriter, errMsg string, code int) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": errMsg})
}
