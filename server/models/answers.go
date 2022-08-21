package models

import (
	"errors"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Answer struct {
	ID         uint   `json:"id"`
	QuestionID uint   `json:"question_id"`
	Content    string `json:"content"`
	IsProper   bool   `json:"is_proper"`
}

type PublicAnswer struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
}

func IsAnswerFromUsersActiveTest(userId uint, answerId uint) (bool, error) {
	var counter uint
	err := GetDB().QueryRow("SELECT COUNT(*) FROM `answers` INNER JOIN questions ON answers.question_id = questions.id INNER JOIN tests ON questions.test_id = tests.id INNER JOIN results ON results.test_id = tests.id WHERE results.finish_datetime IS NULL AND results.start_datetime < NOW() AND results.user_id = ? AND answers.id = ?", userId, answerId).Scan(&counter)
	if err != nil {
		log.Error(fmt.Sprintf("models.IsAnswerOfUsersActiveTest/ %s", err.Error()))
		return false, errors.New(http.StatusText(http.StatusInternalServerError))
	}

	if counter == 0 {
		return false, nil
	}

	return true, nil
}

func DoesAnswerExistInUsersActiveTest(userId uint, answerId uint) (bool, error) {
	var counter uint
	err := GetDB().QueryRow("SELECT COUNT(*) FROM `users_answers` INNER JOIN answers ON answers.id = users_answers.answer_id INNER JOIN questions ON answers.question_id = questions.id INNER JOIN tests ON questions.test_id = tests.id INNER JOIN results ON results.test_id = tests.id WHERE results.finish_datetime IS NULL AND results.start_datetime < NOW() AND results.user_id = ? AND answers.id = ? AND users_answers.created_at > results.start_datetime", userId, answerId).Scan(&counter)
	if err != nil {
		log.Error(fmt.Sprintf("models.DoesAnswerExistInUsersActiveTest/ %s", err.Error()))
		return false, errors.New(http.StatusText(http.StatusInternalServerError))
	}

	if counter == 0 {
		return false, nil
	}

	return true, nil
}
