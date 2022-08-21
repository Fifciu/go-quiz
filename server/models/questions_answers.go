package models

import (
	"errors"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type QuestionWithAnswers struct {
	PublicQuestion
	Answers []*PublicAnswer `json:"answers"`
}

func GetQuestionsAndAnswers(testId uint) ([]*QuestionWithAnswers, error) {
	res, err := GetDB().Query("SELECT questions.id, questions.content, answers.id, answers.content, answers.question_id FROM questions INNER JOIN answers ON answers.question_id = questions.id WHERE test_id = ?", testId)
	if err != nil {
		log.Error(fmt.Sprintf("models.GetQuestionsAndAnswers/ %s", err.Error()))
		return nil, errors.New(http.StatusText(http.StatusInternalServerError))
	}

	questionsMap := make(map[uint]*QuestionWithAnswers)

	for res.Next() {
		question := &PublicQuestion{}
		answer := &Answer{}
		err := res.Scan(&question.ID, &question.Content, &answer.ID, &answer.Content, &answer.QuestionID)
		if err != nil {
			log.Error(fmt.Sprintf("models.GetQuestionsAndAnswers/ %s", err.Error()))
			return nil, errors.New(http.StatusText(http.StatusInternalServerError))
		}

		publicAnswer := &PublicAnswer{ID: answer.ID, Content: answer.Content}
		if _, ok := questionsMap[question.ID]; !ok {
			questionsMap[question.ID] = &QuestionWithAnswers{
				PublicQuestion: PublicQuestion{
					ID:      question.ID,
					Content: question.Content,
				},
				Answers: []*PublicAnswer{publicAnswer}}
		} else {
			questionsMap[question.ID].Answers = append(questionsMap[question.ID].Answers, publicAnswer)
		}
	}

	response := make([]*QuestionWithAnswers, len(questionsMap))
	i := 0
	for _, value := range questionsMap {
		response[i] = value
		i++
	}

	return response, nil
}
