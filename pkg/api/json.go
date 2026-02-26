package api

import (
	"encoding/json"
	"net/http"
)

func JsonError(w http.ResponseWriter, errorMessage string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": errorMessage,
	})
}

func JsonResponse(w http.ResponseWriter, message interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		// Handle potential encoding errors
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
