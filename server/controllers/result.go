package controllers

import (
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
	// POST /results/:test-id/start AUTH aka Start test
	// POST /results/:test-id/finish finishes test if every question is answered
	// GET /results/:test-id Gives results if test is finished
	baseUrl := r.URL.Path[len("/results/"):]
	parts := strings.Split(baseUrl, "/")

	testId, err := strconv.Atoi(parts[0])
	if err != nil || testId < 1 {
		utils.JsonErrorResponse(w, http.StatusNotFound, "Wrong ID")
		return
	}
	values := r.Context().Value("user").(*utils.Claims)
	userId := values.ID
	if r.Method == "POST" && len(parts) == 2 {
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
			utils.JsonResponse(w, http.StatusCreated, newResult)
			return
		} else if parts[1] == "finish" {
			// call hasEveryAnswer()
			// call finish(userId, testId, resultId)
		}
	} else if r.Method == "GET" && len(parts) == 1 {
		// call getResults(userId, testId)
	}

	utils.JsonErrorResponse(w, http.StatusNotFound, "Page not found")
	return
}
