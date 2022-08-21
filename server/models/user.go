package models

import (
	"fmt"
	"errors"
	"net/http"
	log "github.com/sirupsen/logrus"
	mysql "github.com/go-sql-driver/mysql"
)

type User struct {
	ID	uint	`json:"id"`
	Fullname string `json:"fullname"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserPublic struct {
	ID	uint	`json:"id"`
	Fullname string `json:"fullname"`
	Email string `json:"email"`
}

func CreateUser (fullname string, email string, password string) (*User, error) {
	query, err := GetDB().Prepare("INSERT INTO users(fullname, email, password) VALUES (?, ?, ?)");
	if err != nil {
		log.Error(fmt.Sprintf("models.CreateUser/ %s", err.Error()));
		return nil, errors.New(http.StatusText(http.StatusInternalServerError));
	}

	res, err := query.Exec(fullname, email, password);
	if err != nil {
		me, ok := err.(*mysql.MySQLError);
		if !ok {
			log.Error(fmt.Sprintf("models.CreateUser/ %s", err.Error()));
			return nil, errors.New(http.StatusText(http.StatusInternalServerError));
		}
		if me.Number == 1062 {
			log.Debug(fmt.Sprintf("models.CreateUser/ %s", err.Error()));
			return nil, errors.New("Email already used");
		}
		log.Error(fmt.Sprintf("models.CreateUser/ %s", err.Error()));
		return nil, errors.New(http.StatusText(http.StatusInternalServerError));
	}

	lastId, err := res.LastInsertId();
	if err != nil {
		log.Error(fmt.Sprintf("models.CreateUser/ %s", err.Error()));
		return nil, errors.New(http.StatusText(http.StatusInternalServerError));
	}

	return &User{
		ID: uint(lastId),
		Fullname: fullname,
		Email: email,
		Password: password,
	}, nil
}

func GetUserByEmail (email string) (*User, error) {
	user := &User{};
	err := GetDB().QueryRow("SELECT * FROM users WHERE email = ?", email).Scan(&user.ID, &user.Fullname, &user.Email, &user.Password);
	if err != nil {
		log.Error(fmt.Sprintf("models.GetUserByEmail/ %s", err.Error()));
		return nil, errors.New(http.StatusText(http.StatusInternalServerError));
	}

	return user, nil;
}
