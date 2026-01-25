package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// respondWithJSON writes a JSON response with the given status code
func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if payload != nil {
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	}
}

// respondWithError writes an error response with the given status code and message
func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	respondWithJSON(w, statusCode, map[string]interface{}{
		"success": false,
		"error":   message,
	})
}

// extractIDFromPath extracts the ID from the URL path
// Example: /categories/123 -> 123
func extractIDFromPath(r *http.Request, prefix string) (int, error) {
	path := r.URL.Path
	idStr := strings.TrimPrefix(path, prefix)

	// Remove trailing slash if present
	idStr = strings.TrimSuffix(idStr, "/")

	// Parse the ID
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// isNotFoundError checks if the error message indicates a not found error
func isNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	errMsg := err.Error()
	return strings.Contains(errMsg, "not found") ||
		strings.Contains(errMsg, "does not exist")
}
