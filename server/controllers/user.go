package controllers

import (
	"net/http"
	"encoding/json"
	"log"
	models "github.com/Fifciu/go-quiz/server/models"
)

type UserCreate struct {
	Fullname string `json:"fullname"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func CreateUser (w http.ResponseWriter, r *http.Request) {
	// curl -XPOST -d "{\"fullname\":\"John Doe\", \"email\":\"johndoe@gmail.com\", \"password\":\"zaq1@WSX\"}" http://localhost:8090/user
	d := json.NewDecoder(r.Body);
	d.DisallowUnknownFields();

	var body UserCreate;
	err := d.Decode(&body);
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest);
		return;
	}

	if body.Fullname == "" {
		http.Error(w, "Missing field 'fullname' from JSON Object", http.StatusBadRequest);
		return;
	}

	if body.Email == "" {
		http.Error(w, "Missing field 'email' from JSON Object", http.StatusBadRequest);
		return;
	}

	if body.Password == "" {
		http.Error(w, "Missing field 'password' from JSON Object", http.StatusBadRequest);
		return;
	}

	if d.More() {
		http.Error(w, "Extraneous data after JSON object", http.StatusBadRequest);
		return;
	}

	log.Println(body.Fullname);
	log.Println(body.Email);
	log.Println(body.Password);

	user := models.CreateUser(body.Fullname, body.Email, body.Password);

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user);
	return;
}
