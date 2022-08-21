package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/Fifciu/go-quiz/server/models"
	"github.com/Fifciu/go-quiz/server/utils"
)

func GetTests(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		utils.JsonErrorResponse(w, http.StatusNotFound, "Page not found")
		return
	}
	tests, err := models.GetTests()
	if err != nil {
		log.Error(fmt.Sprintf("controllers.GetTests / %s", err.Error()))
		utils.JsonErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JsonResponse(w, 200, tests)
}

func GetTestsQuestionsAndAnswers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		utils.JsonErrorResponse(w, http.StatusNotFound, "Page not found")
		return
	}

	// [ ] GET /tests/:test_id/questions/answers AUTH; remove is_proper!
	baseUrl := r.URL.Path[len("/tests/"):]
	parts := strings.Split(baseUrl, "/")

	if len(parts) != 3 {
		utils.JsonErrorResponse(w, http.StatusNotFound, "Page not found")
		return
	}

	testId, err := strconv.Atoi(parts[0])
	if err != nil || testId < 1 {
		utils.JsonErrorResponse(w, http.StatusNotFound, "Wrong Test ID")
		return
	}

	values := r.Context().Value("user").(*utils.Claims)
	userId := values.ID
	_, err = models.GetResultByTestIdAndUserId(uint(testId), userId)
	if err != nil {
		// User didn't start quiz! 403
		utils.JsonErrorResponse(w, http.StatusForbidden, "User didn't start a test")
		return
	}

	response, err := models.GetQuestionsAndAnswers(uint(testId))
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}
	utils.JsonResponse(w, http.StatusOK, response)
	return
}
