package handler

import (
	"encoding/json"
	"github.com/IamSarojtmg/quod-project/pkg/model"
	"html/template"
	"log"
	"net/http"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now() // Record start time for logging

	data, delta := model.GenerateTimeSeries()
	log.Println("Rendering time series data")
	tmpl := template.Must(template.ParseFiles("web/templates/index.html"))

	err := tmpl.Execute(w, struct {
		Data    []model.TimeSeries
		Start   int
		End     int
		Delta   int
		Current int
	}{
		Data:    data,
		Start:   data[0].Value,
		End:     data[len(data)-1].Value,
		Delta:   delta,
		Current: data[len(data)-1].Value,
	})
	if err != nil {
		log.Printf("ERROR: Failed to render template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Printf("INFO: %s %s [200 OK] - %v", r.Method, r.URL.Path, time.Since(start))
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now() // Record start time for logging
	data, _ := model.GenerateTimeSeries()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		log.Printf("ERROR: Failed to encode JSON response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Printf("INFO: %s %s [200 OK] - %v", r.Method, r.URL.Path, time.Since(start))
}
