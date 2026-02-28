package api

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func JsonError(w http.ResponseWriter, errorMessage string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(map[string]interface{}{
		"error": errorMessage,
	})
	if err != nil {
		log.Error("Error while serializing data to JSON:", err)
	}
}

func JsonResponse(w http.ResponseWriter, message interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		log.Error("Error while serializing data to JSON:", err)
	}
}
