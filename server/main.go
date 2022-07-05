package main

import (
	"log"
	"net/http"
	"os"
	"fmt"
	controllers "github.com/Fifciu/go-quiz/server/controllers"
)

func main() {
	// pwaProtocol := os.Getenv("pwa_protocol");
	// pwaHost := os.Getenv("pwa_host");
	// pwaPort := os.Getenv("pwa_port");
	// pwaHost := os.Getenv("pwa_host");
	// http.HandleFunc("/auth", Auth);
	// http.HandleFunc("/refresh", Refresh);
	http.HandleFunc("/login", controllers.LoginUser);
	http.HandleFunc("/register", controllers.CreateUser);

	apiProtocol := os.Getenv("api_protocol");
	apiHost := os.Getenv("api_host");
	apiPort := os.Getenv("api_port");
	if apiPort == "" {
		apiPort = "8090";
	}
	fmt.Println(fmt.Sprintf("It works on %s://%s:%s address", apiProtocol, apiHost, apiPort));

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", apiPort), nil));
}

