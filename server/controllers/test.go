package controllers

import (
	"net/http"

	"fmt"

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
