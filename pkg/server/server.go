package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yazilimcimekani/mdoctor-server/pkg/controllers"
)

func Start(port string) {
	var startMessage string = fmt.Sprintf("Server started on http://localhost:%s", port)

	// Define the routes
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/markdown", controllers.Markdown())

	// Start the server with logging
	log.Println(startMessage)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
