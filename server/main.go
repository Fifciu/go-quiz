package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	controllers "github.com/Fifciu/go-quiz/server/controllers"
	middlewares "github.com/Fifciu/go-quiz/server/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/users", func(r chi.Router) {
		r.Post("/register", controllers.CreateUser)
		r.Post("/login", controllers.LoginUser)
		r.Post("/refresh", controllers.RefreshToken)
		r.With(middlewares.Authenticated).Post("/me", controllers.UserMe)
	})

	r.Route("/tests", func(r chi.Router) {
		r.Use(middlewares.Authenticated)
		r.Get("/", controllers.GetTests)
		r.Get("/results", controllers.GetEachTestResults)
		r.Get("/{testId}/questions/answers", controllers.GetTestsQuestionsAndAnswers)
	})

	r.With(middlewares.Authenticated).Put("/answers/{answerId}", controllers.PutAnswer)

	r.Route("/results", func(r chi.Router) {
		r.Use(middlewares.Authenticated)
		r.Post("/{testId}/start", controllers.StartTest)
		r.Post("/{resultId}/finish", controllers.FinishTest)
		r.Get("/{resultId}", controllers.GetResults)
	})

	// Each test has: ID, title, image, description
	// Each question has: ID, test_id, content
	// Each answer has: ID, question_id, content, is_proper
	// Each user_answer has: ID, user_id, answer_id, datetime
	// Each result has ID, user_id, test_id, start_datetime, finish_datetime,

	apiProtocol := os.Getenv("api_protocol")
	apiHost := os.Getenv("api_host")
	apiPort := os.Getenv("api_port")
	if apiPort == "" {
		apiPort = "8090"
	}
	fmt.Println(fmt.Sprintf("It works on %s://%s:%s address", apiProtocol, apiHost, apiPort))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", apiPort), r))
}
