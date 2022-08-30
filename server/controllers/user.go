package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	models "github.com/Fifciu/go-quiz/server/models"
	utils "github.com/Fifciu/go-quiz/server/utils"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserCreate struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignedUser struct {
	Token          string    `json:"token"`
	ExpirationTime time.Time `json:"expirationTime"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.JsonErrorResponse(w, http.StatusNotFound, "Page not found")
		return
	}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	var body UserCreate
	err := d.Decode(&body)
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if body.Fullname == "" {
		utils.JsonErrorResponse(w, http.StatusBadRequest, "Missing field 'fullname'")
		return
	}

	if body.Email == "" {
		utils.JsonErrorResponse(w, http.StatusBadRequest, "Missing field 'email'")
		return
	}

	if body.Password == "" {
		utils.JsonErrorResponse(w, http.StatusBadRequest, "Missing field 'password'")
		return
	}

	if d.More() {
		utils.JsonErrorResponse(w, http.StatusBadRequest, "Extraneous data after JSON object")
		return
	}

	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	user, err := models.CreateUser(body.Fullname, body.Email, hashedPassword)
	if err != nil {
		if err.Error() == http.StatusText(http.StatusInternalServerError) {
			utils.JsonErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		utils.JsonErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	token, expirationTime, err := utils.GenerateJwtToken(models.UserPublic{
		ID:       user.ID,
		Fullname: user.Fullname,
		Email:    user.Email,
	})

	if err != nil {
		log.Error(fmt.Sprintf("controllers.CreateUser / %s", err.Error()))
		utils.JsonErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JsonResponse(w, http.StatusCreated, &SignedUser{
		Token:          token,
		ExpirationTime: expirationTime,
	})
	return
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	var body UserLogin
	err := d.Decode(&body)
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if body.Email == "" {
		utils.JsonErrorResponse(w, http.StatusBadRequest, "Missing field 'email'")
		return
	}

	if body.Password == "" {
		utils.JsonErrorResponse(w, http.StatusBadRequest, "Missing field 'password'")
		return
	}

	if d.More() {
		utils.JsonErrorResponse(w, http.StatusBadRequest, "Extraneous data after JSON object")
		return
	}

	user, err := models.GetUserByEmail(body.Email)
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusBadRequest, "Wrong email or password")
		return
	}

	token, expirationTime, err := utils.GenerateJwtToken(models.UserPublic{
		ID:       user.ID,
		Fullname: user.Fullname,
		Email:    user.Email,
	})

	if err != nil {
		log.Error(fmt.Sprintf("controllers.LoginUser / %s", err.Error()))
		utils.JsonErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JsonResponse(w, http.StatusOK, &SignedUser{
		Token:          token,
		ExpirationTime: expirationTime,
	})
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	cookieKey := os.Getenv("cookie_token_key")
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
		fmt.Println(err.Error())
		utils.JsonErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if !tkn.Valid {
		utils.JsonErrorResponse(w, http.StatusUnauthorized, "Invalid token. Unauthorized")
		return
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		utils.JsonErrorResponse(w, http.StatusBadRequest, "It's too early to refresh the token")
		return
	}

	jwtTtlString := os.Getenv("jwt_ttl")
	jwtTtl, err := strconv.Atoi(jwtTtlString)
	if err != nil {
		log.Error(fmt.Sprintf("controllers.RefreshToken / %s", err.Error()))
		utils.JsonErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	expirationTime := time.Now().Add(time.Duration(jwtTtl) * time.Second)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		log.Error(fmt.Sprintf("controllers.RefreshToken / %s", err.Error()))
		utils.JsonErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JsonResponse(w, http.StatusOK, &SignedUser{
		Token:          tokenString,
		ExpirationTime: expirationTime,
	})
}

func UserMe(w http.ResponseWriter, r *http.Request) {
	values := r.Context().Value("user").(*utils.Claims)
	utils.JsonResponse(w, http.StatusOK, &models.UserPublic{
		ID:       values.ID,
		Fullname: values.Fullname,
		Email:    values.Email,
	})
}
