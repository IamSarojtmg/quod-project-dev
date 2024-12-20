package handler

import (
	"github.com/IamSarojtmg/quod-project/pkg/model"
	"html/template"
	"log/slog"
	"net/http"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now() // Record start time for logging

	data, delta := model.GenerateTimeSeries()
	slog.Info("Rendering time series data") // Use slog for simple informational logs

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
		// Structured error logging with additional context
		slog.Error("Failed to render template",
			"error", err,
			"path", r.URL.Path,
			"method", r.Method,
		)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Log successful request processing with additional details
	slog.Info("Request handled successfully",
		"method", r.Method,
		"path", r.URL.Path,
		"status", http.StatusOK,
		"duration", time.Since(start),
	)
}

