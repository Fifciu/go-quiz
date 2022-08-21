package models

import (
	"errors"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Test struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
}

func GetTests() ([]*Test, error) {
	tests := []*Test{}
	res, err := GetDB().Query("SELECT * FROM tests")
	if err != nil {
		log.Error(fmt.Sprintf("models.GetTests/ %s", err.Error()))
		return nil, errors.New(http.StatusText(http.StatusInternalServerError))
	}

	for res.Next() {
		test := &Test{}
		err := res.Scan(&test.ID, &test.Title, &test.Description, &test.ImageUrl)
		if err != nil {
			log.Error(fmt.Sprintf("models.GetTests/ %s", err.Error()))
			return nil, errors.New(http.StatusText(http.StatusInternalServerError))
		}
		tests = append(tests, test)
	}

	return tests, nil
}
