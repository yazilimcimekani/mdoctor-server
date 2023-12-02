package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Define the handler for the /public/ route
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	// Define the handler for the markdown route
	http.HandleFunc("/markdown", markdown())

	// Get the port from the environment or use the default port of 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server on given port
	startMessage := fmt.Sprintf("Server started on http://localhost:%s", port)
	log.Println(startMessage)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func loadFile(htmlFilePath string) string {
	htmlContent, err := os.ReadFile(htmlFilePath)

	if err != nil {
		log.Fatal(err)
	}

	return string(htmlContent)
}

func mdToHtml(md string) string {
	if md == "# Hello World" {
		md = fmt.Sprintf(`<h1>Hello World</h1>`)
	}
	return md
}

func markdown() func(w http.ResponseWriter, r *http.Request) {
	const mdTemplatePath = "views/md.html"
	var html string = loadFile(mdTemplatePath)
	fullHTML := fmt.Sprintf(html, mdToHtml("# Hello World"))

	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fullHTML))
	}
}
