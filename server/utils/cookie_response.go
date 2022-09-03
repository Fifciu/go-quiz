package utils

import (
	"net/http"
	"os"
	"time"
)

func CookieResponse(w http.ResponseWriter, token string, expirationTime time.Time) {
	w.Header().Set("Content-Type", "application/json")
	cookieKey := os.Getenv("cookie_token_key")
	http.SetCookie(w, &http.Cookie{
		Name:    cookieKey,
		Value:   token,
		Expires: expirationTime,
		Path:    "/",
	})
}
