package main

import (
	// "fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
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

	// Handle CORS
	router.Use(cors.Handler(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedOrigins: []string{"https://*", "http://*"},
		ExposedHeaders: []string{"Link"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	// Router to healthz
	v1Router := chi.NewRouter()
	v1Router.HandleFunc("/healthz", handlerReadiness)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:	":" + portNum,
	}

	// Log server start
	log.Println("Server starting on port:", portNum)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
