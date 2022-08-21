package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	utils "github.com/Fifciu/go-quiz/server/utils"
	"github.com/dgrijalva/jwt-go"
)

func Authenticated(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	cookieKey := os.Getenv("cookie_token_key")
	if cookieKey == "" {
		log.Fatal("Couldn't find value of 'cookie_token_key' environment variable!")
		panic("Couldn't find value of 'cookie_token_key' environment variable!")
	}
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie(cookieKey)
		if err != nil {
			if err == http.ErrNoCookie {
				utils.JsonErrorResponse(w, http.StatusUnauthorized, "Authorization cookie not sent")
				return
			}
			utils.JsonErrorResponse(w, http.StatusBadRequest, "Bad request")
			return
		}
		tokenFromCookie := c.Value

		claims := &utils.Claims{}
		jwtKey := os.Getenv("jwt_key")

		tkn, err := jwt.ParseWithClaims(tokenFromCookie, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				utils.JsonErrorResponse(w, http.StatusUnauthorized, "Invalid signature. Unauthorized")
				return
			}
			utils.JsonErrorResponse(w, http.StatusBadRequest, err.Error()) // Is it safe to pass it??
			return
		}

		if !tkn.Valid {
			utils.JsonErrorResponse(w, http.StatusUnauthorized, "Invalid token. Unauthorized")
			return
		}

		fmt.Println("Authenticated middleware")
		ctx := context.WithValue(r.Context(), "user", claims)
		next(w, r.WithContext(ctx))
	}
}
