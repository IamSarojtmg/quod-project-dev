package main

import (
	"encoding/json"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

type TimeSeries struct {
	Time  string
	Value int
}

func generateTimeSeries() ([]TimeSeries, int) {
	data := make([]TimeSeries, 10)
	var startValue int
	currentValue := startValue

	for i := 0; i < 10; i++ {
		timestamp := time.Now().Add(time.Duration(i) * time.Hour).Format("15:00")
		currentValue += rand.Intn(5) - 2
		data[i] = TimeSeries{
			Time:  timestamp,
			Value: currentValue,
		}
	}
	delta := data[len(data)-1].Value - data[0].Value
	return data, delta
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data, delta := generateTimeSeries()
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, struct {
		Data    []TimeSeries
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
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := generateTimeSeries()
	w.Header().Set("Content-Type", " application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/time-series", apiHandler)

	println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
