package controllers

import (
	"net/http"
	"strconv"

	"fmt"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"

	"github.com/Fifciu/go-quiz/server/models"
	"github.com/Fifciu/go-quiz/server/utils"
)

func GetTests(w http.ResponseWriter, r *http.Request) {
	tests, err := models.GetTests()
	if err != nil {
		log.Error(fmt.Sprintf("controllers.GetTests / %s", err.Error()))
		utils.JsonErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JsonResponse(w, 200, tests)
}

func GetTestsQuestionsAndAnswers(w http.ResponseWriter, r *http.Request) {
	values := r.Context().Value("user").(*utils.Claims)
	userId := values.ID

	testId, err := strconv.Atoi(chi.URLParam(r, "testId"))
	if err != nil || testId < 1 {
		utils.JsonErrorResponse(w, http.StatusNotFound, "Wrong Test ID")
		return
	}

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

func GetEachTestResults(w http.ResponseWriter, r *http.Request) {
	values := r.Context().Value("user").(*utils.Claims)
	userId := values.ID

	stats, err := models.GetUserEveryResults(userId)
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}
	utils.JsonResponse(w, http.StatusOK, stats)
	return
}
