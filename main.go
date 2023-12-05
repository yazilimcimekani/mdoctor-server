package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/yazilimcimekani/mdoctor-server/pkg/env"
	"github.com/yazilimcimekani/mdoctor-server/pkg/server"
)

func main() {
	// Load the .env file
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	// If the port is not set, use the default port
	port := env.GetEnv("PORT", "8080")

	server.Start(port)
}
