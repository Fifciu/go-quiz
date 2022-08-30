// Utility package jou
package utils

import (
	"os"
	"strconv"
	"time"

	models "github.com/Fifciu/go-quiz/server/models"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	ID       uint   `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// Generates JWT Token
func GenerateJwtToken(user models.UserPublic) (string, time.Time, error) {
	// JWT
	jwtTtlString := os.Getenv("jwt_ttl")
	jwtTtl, err := strconv.Atoi(jwtTtlString)
	if err != nil {
		return "", time.Time{}, err
	}
	expirationTime := time.Now().Add(time.Duration(jwtTtl) * time.Second)
	claims := &Claims{
		ID:       user.ID,
		Fullname: user.Fullname,
		Email:    user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtKey := os.Getenv("jwt_key")
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expirationTime, nil
}
