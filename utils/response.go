package utils

import (
	"encoding/json"
	"net/http"
)

// Respond sends a JSON response
func Respond(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
