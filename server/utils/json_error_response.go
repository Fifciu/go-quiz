package utils

import (
	"net/http"
	"encoding/json"
)

func JsonErrorResponse(w http.ResponseWriter, statusCode int, errMsg string) {
	w.Header().Set("Content-Type", "application/json");
	w.WriteHeader(statusCode);
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": errMsg,
	});
}
