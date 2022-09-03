package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	controllers "github.com/Fifciu/go-quiz/server/controllers"
	middlewares "github.com/Fifciu/go-quiz/server/middlewares"
	"github.com/Fifciu/go-quiz/server/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{utils.BuildClientBaseUrl()},
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
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

	apiProtocol := os.Getenv("api_protocol")
	apiHost := os.Getenv("api_host")
	apiPort := os.Getenv("api_port")
	if apiPort == "" {
		apiPort = "8090"
	}
	fmt.Println(fmt.Sprintf("It works on %s://%s:%s address", apiProtocol, apiHost, apiPort))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", apiPort), r))
}
