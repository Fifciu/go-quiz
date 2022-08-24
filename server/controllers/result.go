package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Fifciu/go-quiz/server/models"
	"github.com/Fifciu/go-quiz/server/utils"
)

func ResultsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" && r.Method != "POST" {
		utils.JsonErrorResponse(w, http.StatusNotFound, "Page not found")
		return
	}
	// TODO: Maybe the first one shouldn't start with /results?
	// POST /results/:test-id/start AUTH aka Start test
	// POST /results/:result-id/finish finishes test if every question is answered
	// GET /results/:test-id Gives results if test is finished
	baseUrl := r.URL.Path[len("/results/"):]
	parts := strings.Split(baseUrl, "/")

	testId, err := strconv.Atoi(parts[0])
	if err != nil || testId < 1 {
		utils.JsonErrorResponse(w, http.StatusNotFound, "Wrong Test ID")
		return
	}
	values := r.Context().Value("user").(*utils.Claims)
	userId := values.ID
	if r.Method == "POST" && len(parts) == 2 {
		// TODO: Switch
		if parts[1] == "start" {
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
			// Lack of start_datetime
			utils.JsonResponse(w, http.StatusCreated, newResult)
			return
		} else if parts[1] == "finish" {
			// TODO: Guard, not exising ID (can do it based on 0 value inside )
			has, err := models.HasEveryAnswer(userId, uint(testId)) // in this case its resultId
			if err != nil {
				if err.Error() == http.StatusText(http.StatusNotFound) {
					errMessage := fmt.Sprintf("Couldn't find active test session with provided ID (%d) for this current user (%d).", testId, userId)
					utils.JsonErrorResponse(w, http.StatusNotFound, errMessage)
					return
				} else {
					utils.JsonErrorResponse(w, http.StatusInternalServerError, err.Error())
					return
				}
			}

			if has {
				err = models.ResultFinish(uint(testId), userId)
				if err != nil {
					utils.JsonErrorResponse(w, http.StatusNotFound, err.Error())
					return
				}
				stats, err := models.GetUserResults(uint(testId), userId)
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
	} else if r.Method == "GET" && len(parts) == 1 {
		stats, err := models.GetUserResults(uint(testId), userId)
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

	utils.JsonErrorResponse(w, http.StatusNotFound, "Page not found")
	return
}
