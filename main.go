package main

import (
	"log"

	"github.com/EmilioCliff/portfolio.git/api"
)

const (
	serverAddress = "0.0.0.0:8080"
)

func main() {
	log.Println("Starting server at port: 8080")
	if err := api.StartServer(serverAddress); err != nil {
		log.Fatal("Couldn't start server: %w", err)
	}
}
