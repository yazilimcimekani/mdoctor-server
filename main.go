package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/yazilimcimekani/mdoctor-server/pkg/controllers"
	"github.com/yazilimcimekani/mdoctor-server/pkg/env"
)

func main() {
	// Load the .env file
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
	port := env.GetEnv("PORT", "8080")

	// Define the handlers
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/markdown", controllers.Markdown())

	// Start the server on given port
	startMessage := fmt.Sprintf("Server started on http://localhost:%s", port)
	log.Println(startMessage)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
