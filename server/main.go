package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	controllers "github.com/Fifciu/go-quiz/server/controllers"
	middlewares "github.com/Fifciu/go-quiz/server/middlewares"
)

func main() {
	http.HandleFunc("/register", controllers.CreateUser)
	http.HandleFunc("/login", controllers.LoginUser)
	http.HandleFunc("/me", middlewares.Authenticated(controllers.UserMe))

	http.HandleFunc("/tests", middlewares.Authenticated(controllers.GetTests))
	http.HandleFunc("/tests/", middlewares.Authenticated(controllers.GetTestsQuestionsAndAnswers))
	http.HandleFunc("/results/", middlewares.Authenticated(controllers.ResultsHandler)) // DAMN
	http.HandleFunc("/answers/", middlewares.Authenticated(controllers.PutAnswer))

	// Each test has: ID, title, image, description
	// Each question has: ID, test_id, content
	// Each answer has: ID, question_id, content, is_proper
	// Each user_answer has: ID, user_id, answer_id, datetime
	// Each result has ID, user_id, test_id, start_datetime, finish_datetime,

	// TODO: Export new DB

	// Endpoints
	// [X] GET /tests AUTH
	// GET /tests/results AUTH-PER-USER
	// [x] GET /tests/:test_id/questions/answers AUTH; remove is_proper, prohibited if user didn't start quiz
	// [x] PUT /answers/:answer-id AUTH; adds user_answer
	// [x] POST /results/:test-id/start AUTH aka Start test
	// POST /results/:test-id/finish finishes test if every question is answered
	// GET /results/:test-id Gives results if test is finished

	apiProtocol := os.Getenv("api_protocol")
	apiHost := os.Getenv("api_host")
	apiPort := os.Getenv("api_port")
	if apiPort == "" {
		apiPort = "8090"
	}
	fmt.Println(fmt.Sprintf("It works on %s://%s:%s address", apiProtocol, apiHost, apiPort))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", apiPort), nil))
}
