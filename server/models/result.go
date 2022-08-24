package models

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	mysql "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

type DateType time.Time

func (d DateType) String() string {
	return time.Time(d).String()
}

type Result struct {
	ID             uint     `json:"id"`
	TestID         uint     `json:"test_id"`
	UserID         uint     `json:"user_id"`
	StartDatetime  DateType `json:"start_datetime"`
	FinishDatetime DateType `json:"finish_datetime"`
}

type QuestionAndUserAnswer struct {
	TestId   uint   `json:"test_id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
	IsProper bool   `json:"is_proper"`
}

func ResultStart(testId uint, userId uint) (*Result, error) {
	query, err := GetDB().Prepare("INSERT INTO results(test_id, user_id) VALUES (?, ?)")
	if err != nil {
		log.Error(fmt.Sprintf("models.ResultStart/ %s", err.Error()))
		return nil, errors.New(http.StatusText(http.StatusInternalServerError))
	}

	res, err := query.Exec(testId, userId)
	if err != nil {
		me, ok := err.(*mysql.MySQLError)
		if !ok {
			log.Error(fmt.Sprintf("models.ResultStart/ %s", err.Error()))
			return nil, errors.New(http.StatusText(http.StatusInternalServerError))
		}
		if me.Number == 1452 {
			log.Error(fmt.Sprintf("models.ResultStart/ %s", err.Error()))
			return nil, errors.New("Wrong Test IDs")
		}
		log.Error(fmt.Sprintf("models.ResultStart/ %s", err.Error()))
		return nil, errors.New(http.StatusText(http.StatusInternalServerError))
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Error(fmt.Sprintf("models.ResultStart/ %s", err.Error()))
		return nil, errors.New(http.StatusText(http.StatusInternalServerError))
	}

	return &Result{
		ID:            uint(lastId),
		TestID:        testId,
		UserID:        userId,
		StartDatetime: DateType(time.Now()),
	}, nil
}

func ResultFinish(resultId uint, userId uint) error {
	_, err := GetDB().Query("UPDATE results SET finish_datetime = ? WHERE id = ? AND user_id = ? AND finish_datetime IS NULL", time.Now().Format("2006-01-02 15:04:05"), resultId, userId)
	if err != nil {
		log.Error(fmt.Sprintf("models.ResultFinish/ %s", err.Error()))
		return errors.New(http.StatusText(http.StatusInternalServerError))
	}
	return nil
}

func GetUserResults(resultId uint, userId uint) ([]*QuestionAndUserAnswer, error) {
	res, err := GetDB().Query("SELECT results.test_id, questions.content, answers.is_proper, answers.content FROM `results` INNER JOIN questions ON questions.test_id = results.test_id INNER JOIN answers ON answers.question_id = questions.id INNER JOIN users_answers ON users_answers.answer_id = answers.id WHERE results.id = ? AND results.user_id = ? AND users_answers.created_at < results.finish_datetime AND users_answers.created_at > results.start_datetime;", resultId, userId)
	if err != nil {
		log.Error(fmt.Sprintf("models.GetQuestionsAndAnswers/ %s", err.Error()))
		return nil, errors.New(http.StatusText(http.StatusInternalServerError))
	}

	response := make([]*QuestionAndUserAnswer, 0)
	for res.Next() {
		unit := &QuestionAndUserAnswer{}
		res.Scan(&unit.TestId, &unit.Question, &unit.IsProper, &unit.Answer)
		response = append(response, unit)
	}
	return response, nil
}

func GetUserEveryResults(userId uint) ([]*QuestionAndUserAnswer, error) {
	// TODO: Definitely I have to do smth with this one
	tooLongQuery := `SELECT a.test_id, a.content, a.answerContent, a.is_proper FROM (SELECT results.user_id, results.test_id, questions.id as qId, questions.content, answers.is_proper, answers.content as answerContent, answers.id, results.finish_datetime,users_answers.created_at
		FROM results
		INNER JOIN questions ON questions.test_id = results.test_id 
		INNER JOIN answers ON answers.question_id = questions.id 
		INNER JOIN users_answers ON users_answers.answer_id = answers.id 
		WHERE results.user_id = ? AND users_answers.created_at < results.finish_datetime AND users_answers.created_at > results.start_datetime
		ORDER BY results.finish_datetime DESC) as a
		LEFT OUTER JOIN (SELECT results.user_id, results.test_id, questions.id as qId, questions.content, answers.is_proper, answers.content as answerContent, answers.id, results.finish_datetime,users_answers.created_at
		FROM results
		INNER JOIN questions ON questions.test_id = results.test_id 
		INNER JOIN answers ON answers.question_id = questions.id 
		INNER JOIN users_answers ON users_answers.answer_id = answers.id 
		WHERE results.user_id = ? AND users_answers.created_at < results.finish_datetime AND users_answers.created_at > results.start_datetime
		ORDER BY results.finish_datetime DESC
			) as b ON a.qId = b.qId AND a.finish_datetime < b.finish_datetime
			WHERE b.id IS NULL`
	res, err := GetDB().Query(tooLongQuery, userId, userId)
	if err != nil {
		log.Error(fmt.Sprintf("models.GetQuestionsAndAnswers/ %s", err.Error()))
		return nil, errors.New(http.StatusText(http.StatusInternalServerError))
	}

	response := make([]*QuestionAndUserAnswer, 0)
	for res.Next() {
		unit := &QuestionAndUserAnswer{}
		res.Scan(&unit.TestId, &unit.Question, &unit.Answer, &unit.IsProper)
		response = append(response, unit)
	}

	// TODO: Better shape of response interfaces
	return response, nil
}

func GetResultByTestIdAndUserId(testId uint, userId uint) (*Result, error) {
	var finishDatetimeTmp mysql.NullTime
	result := &Result{}
	err := GetDB().QueryRow("SELECT * FROM results WHERE test_id = ? AND user_id = ? and finish_datetime IS NULL", testId, userId).Scan(&result.ID, &result.TestID, &result.UserID, &result.StartDatetime, &finishDatetimeTmp)
	if err != nil {
		log.Error(fmt.Sprintf("models.GetResultByTestIdAndUserId/ %s", err.Error()))
		return nil, errors.New(http.StatusText(http.StatusInternalServerError))
	}
	if finishDatetimeTmp.Valid {
		result.FinishDatetime = DateType(finishDatetimeTmp.Time)
	}

	return result, nil
}
