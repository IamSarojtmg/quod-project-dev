package handler

import (
	"encoding/json"
	"github.com/IamSarojtmg/quod-project/pkg/model"
	"log/slog"
	"net/http"
	"time"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now() // Record start time for logging

	data, _ := model.GenerateTimeSeries()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		// Structured error logging
		slog.Error("Failed to encode JSON response",
			"error", err,
			"path", r.URL.Path,
			"method", r.Method,
		)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Log successful request processing
	slog.Info("API request handled successfully",
		"method", r.Method,
		"path", r.URL.Path,
		"status", http.StatusOK,
		"duration", time.Since(start),
	)
}
