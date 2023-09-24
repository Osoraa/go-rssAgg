package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello World")

	godotenv.Load()

	portNum := os.Getenv("PORT")
	if portNum == "" {
		log.Fatal("PORT not found in your enironment")
	}

	fmt.Println("Port:", portNum)
}
