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

	// Start the server
	log.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
