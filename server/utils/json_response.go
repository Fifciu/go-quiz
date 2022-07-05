package utils

import (
	"net/http"
	"encoding/json"
)

func JsonResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json");
	w.WriteHeader(statusCode);
	json.NewEncoder(w).Encode(payload);
}
