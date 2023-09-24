package main

import (
	// "fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {

	// Load variables from .env file
	godotenv.Load()

	portNum := os.Getenv("PORT")
	if portNum == "" {
		log.Fatal("PORT not found in your enironment")
	}

	// Create a router, server, and then listen
	router := chi.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr:	":" + portNum,
	}

	log.Printf("\nServer starting on port: %s", portNum)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
