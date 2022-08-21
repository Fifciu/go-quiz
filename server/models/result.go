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
