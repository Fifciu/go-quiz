package models

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type UserAnswer struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	AnswerID  uint      `json:"answer_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UserAnswerDraft struct {
	UserID   uint `json:"user_id"`
	AnswerID uint `json:"answer_id"`
}

func CreateUserAnswer(userId uint, answerId uint) (*UserAnswer, error) {
	query, err := GetDB().Prepare("INSERT INTO users_answers(user_id, answer_id) VALUES (?, ?)")
	if err != nil {
		log.Error(fmt.Sprintf("models.CreateUserAnswer/ %s", err.Error()))
		return nil, errors.New(http.StatusText(http.StatusInternalServerError))
	}

	res, err := query.Exec(userId, answerId)
	if err != nil {
		log.Error(fmt.Sprintf("models.CreateUserAnswer/ %s", err.Error()))
		return nil, errors.New(http.StatusText(http.StatusInternalServerError))
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Error(fmt.Sprintf("models.CreateUserAnswer/ %s", err.Error()))
		return nil, errors.New(http.StatusText(http.StatusInternalServerError))
	}

	return &UserAnswer{
		ID:        uint(lastId),
		UserID:    userId,
		AnswerID:  answerId,
		CreatedAt: time.Now(),
	}, nil
}
