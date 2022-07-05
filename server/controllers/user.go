package controllers

import (
	"net/http"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	models "github.com/Fifciu/go-quiz/server/models"
	utils "github.com/Fifciu/go-quiz/server/utils"
)

type UserCreate struct {
	Fullname string `json:"fullname"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserLogin struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserPublic struct {
	ID	uint	`json:"id"`
	Fullname string `json:"fullname"`
	Email string `json:"email"`
}

func CreateUser (w http.ResponseWriter, r *http.Request) {
	// curl -XPOST -d "{\"fullname\":\"John Doe\", \"email\":\"johndoe@gmail.com\", \"password\":\"zaq1@WSX\"}" http://localhost:8090/user
	if r.Method != "POST" {
		http.Error(w, "Page not found", http.StatusNotFound);
		return;
	}
	d := json.NewDecoder(r.Body);
	d.DisallowUnknownFields();

	var body UserCreate;
	err := d.Decode(&body);
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusBadRequest, err.Error());
		return;
	}

	if body.Fullname == "" {
		utils.JsonErrorResponse(w, http.StatusBadRequest, "Missing field 'fullname'");
		return;
	}

	if body.Email == "" {
		utils.JsonErrorResponse(w, http.StatusBadRequest, "Missing field 'email'");
		return;
	}

	if body.Password == "" {
		utils.JsonErrorResponse(w, http.StatusBadRequest, "Missing field 'password'");
		return;
	}

	if d.More() {
		utils.JsonErrorResponse(w, http.StatusBadRequest, "Extraneous data after JSON object");
		return
	}

	hashedPassword, err := utils.HashPassword(body.Password);
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError));
	}

	user, err := models.CreateUser(body.Fullname, body.Email, hashedPassword);
	if err != nil {
		if err.Error() == http.StatusText(http.StatusInternalServerError) {
			utils.JsonErrorResponse(w, http.StatusInternalServerError, err.Error());
			return;
		}
		utils.JsonErrorResponse(w, http.StatusBadRequest, err.Error());
		return;
	}

	utils.JsonResponse(w, http.StatusCreated, user);
	return 
}

func LoginUser (w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Page not found", http.StatusNotFound);
		return;
	}
	d := json.NewDecoder(r.Body);
	d.DisallowUnknownFields();

	var body UserLogin;
	err := d.Decode(&body);
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusBadRequest, err.Error());
		return;
	}

	if body.Email == "" {
		utils.JsonErrorResponse(w, http.StatusBadRequest, "Missing field 'email'");
		return;
	}

	if body.Password == "" {
		utils.JsonErrorResponse(w, http.StatusBadRequest, "Missing field 'password'");
		return;
	}

	if d.More() {
		utils.JsonErrorResponse(w, http.StatusBadRequest, "Extraneous data after JSON object");
		return
	}

	user, err := models.GetUserByEmail(body.Email);
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusInternalServerError, err.Error());
		return;
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password));
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusBadRequest, "Wrong email or password");
		return;
	}

	utils.JsonResponse(w, http.StatusOK, &UserPublic{
		ID: user.ID,
		Fullname: user.Fullname,
		Email: user.Email,
	});
	// TODO: Generate JWT
}
