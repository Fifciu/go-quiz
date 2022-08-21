package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Fifciu/go-quiz/server/models"
	"github.com/Fifciu/go-quiz/server/utils"
)

func PutAnswer(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		utils.JsonErrorResponse(w, http.StatusNotFound, "Page not found")
		return
	}

	baseUrl := r.URL.Path[len("/answers/"):]
	parts := strings.Split(baseUrl, "/")

	answerId, err := strconv.Atoi(parts[0])
	if err != nil || answerId < 1 {
		utils.JsonErrorResponse(w, http.StatusNotFound, "Wrong Answer ID")
		return
	}
	values := r.Context().Value("user").(*utils.Claims)
	userId := values.ID

	// Guard: If already exists for the test. Not older one.
	if is, _ := models.DoesAnswerExistInUsersActiveTest(userId, uint(answerId)); is {
		utils.JsonErrorResponse(w, http.StatusBadRequest, "User already answered to this question")
		return
	}

	// Guard: Is answer of question of test that user is doing? Not older one.
	if is, _ := models.IsAnswerFromUsersActiveTest(userId, uint(answerId)); !is {
		utils.JsonErrorResponse(w, http.StatusBadRequest, "User didn't start or already finished test containing this answer")
		return
	}

	// TODO in v1.1: Update if it already exists
	userAnswer, err := models.CreateUserAnswer(userId, uint(answerId))
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JsonResponse(w, http.StatusCreated, userAnswer)
	return
}
