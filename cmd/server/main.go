package main

import (
	"github.com/IamSarojtmg/quod-project/pkg/handler"
	"log"
	"net/http"
)

func main() {
	staticFileServer := http.FileServer(http.Dir("./web"))
	http.Handle("/web/", http.StripPrefix("/web", staticFileServer))

	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/api/time-series", handler.ApiHandler)

	serverAddress := ":8080"
	log.Printf("INFO: Server is running on http://localhost%s", serverAddress)
	err := http.ListenAndServe(serverAddress, nil)
	if err != nil {
		log.Fatalf("ERROR: Failed to start server: %v", err)
	}
}
