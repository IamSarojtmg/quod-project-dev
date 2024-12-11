package handler

import (
	"log"
	"encoding/json"
	"net/http"
	"github.com/IamSarojtmg/quod-project/pkg/model"
	"html/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data, delta := model.GenerateTimeSeries()
	log.Println("Rendering time series data")
	tmpl := template.Must(template.ParseFiles("web\\templates\\index.html"))
	tmpl.Execute(w, struct {
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
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("API request received: /api/time-series")
	data, _ := model.GenerateTimeSeries()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
