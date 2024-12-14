package main

import (
	"log"
	"net/http"
	"github.com/IamSarojtmg/quod-project/pkg/handler"
)

func main() {
	// Set up routes
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/api/time-series", handler.ApiHandler)

	serverAddress := ":8080"
	log.Printf("INFO: Server is running on http://localhost%s", serverAddress)
	err := http.ListenAndServe(serverAddress, nil)
	if err != nil {
		log.Fatalf("ERROR: Failed to start server: %v", err)
	}
}
