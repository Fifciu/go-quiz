package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Fifciu/go-quiz/server/models"
	"github.com/Fifciu/go-quiz/server/utils"
	"github.com/go-chi/chi/v5"
)

func GetResults(w http.ResponseWriter, r *http.Request) {
	resultId, err := strconv.Atoi(chi.URLParam(r, "resultId"))
	if err != nil || resultId < 1 {
		utils.JsonErrorResponse(w, http.StatusNotFound, "Wrong Result ID")
		return
	}
	values := r.Context().Value("user").(*utils.Claims)
	userId := values.ID
	stats, err := models.GetUserResults(uint(resultId), userId)
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}
	if len(stats) < 1 {
		utils.JsonErrorResponse(w, http.StatusNotFound, "Result with provided ID doesn't exist for the current user")
		return
	}
	utils.JsonResponse(w, http.StatusOK, stats)
	return
}

func StartTest(w http.ResponseWriter, r *http.Request) {
	testId, err := strconv.Atoi(chi.URLParam(r, "testId"))
	if err != nil || testId < 1 {
		utils.JsonErrorResponse(w, http.StatusNotFound, "Wrong Test ID")
		return
	}
	values := r.Context().Value("user").(*utils.Claims)
	userId := values.ID
	alreadyExistingResult, err := models.GetResultByTestIdAndUserId(uint(testId), userId)
	if err == nil {
		// Found and returns
		utils.JsonResponse(w, http.StatusOK, alreadyExistingResult)
		return
	}

	newResult, err := models.ResultStart(uint(testId), userId)
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}
	// TODO: Lack of start_datetime
	utils.JsonResponse(w, http.StatusCreated, newResult)
	return
}

func FinishTest(w http.ResponseWriter, r *http.Request) {
	resultId, err := strconv.Atoi(chi.URLParam(r, "resultId"))
	if err != nil || resultId < 1 {
		utils.JsonErrorResponse(w, http.StatusNotFound, "Wrong Test ID")
		return
	}
	values := r.Context().Value("user").(*utils.Claims)
	userId := values.ID
	// TODO: Guard, not exising ID (can do it based on 0 value inside )
	has, err := models.HasEveryAnswer(userId, uint(resultId)) // in this case its resultId
	if err != nil {
		if err.Error() == http.StatusText(http.StatusNotFound) {
			errMessage := fmt.Sprintf("Couldn't find active test session with provided ID (%d) for this current user (%d).", resultId, userId)
			utils.JsonErrorResponse(w, http.StatusNotFound, errMessage)
			return
		} else {
			utils.JsonErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	if has {
		err = models.ResultFinish(uint(resultId), userId)
		if err != nil {
			utils.JsonErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		stats, err := models.GetUserResults(uint(resultId), userId)
		if err != nil {
			utils.JsonErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		utils.JsonResponse(w, http.StatusOK, stats)
		return
	} else {
		// TODO: Modify every error message to be user friendly for some toast
		utils.JsonErrorResponse(w, http.StatusForbidden, "You didn't response to every answer in the test, cannot finish.")
		return
	}
}
