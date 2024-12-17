package model

import (
	"math/rand"
	"time"
)

// TimeSeries represents a single time series data point.
type TimeSeries struct {
	Time  string
	Value int
}

// GenerateTimeSeries generates time series data with a more meaningful progression.
func GenerateTimeSeries() ([]TimeSeries, int) {
	// Create a local random generator with a seed
	rng := rand.New(rand.NewSource(time.Now().UnixNano())) // Use local random generator
	
	data := make([]TimeSeries, 10)

	// Start with a base value (can be random or fixed)
	startValue := 50 + rng.Intn(20) // Starting value between 50 and 70
	currentValue := startValue

	// Generate time series data with a trend
	for i := 0; i < 10; i++ {
		// Simulate small upward/downward trend with randomness
		trend := 1 - rng.Intn(3)  // Generates -1, 0, or 1 for slight upward/downward trend
		noise := rng.Intn(3) - 1  // Random noise: -1, 0, or 1

		currentValue += trend + noise
		if currentValue < 0 { // Prevent negative values
			currentValue = 0
		}

		// Add the data point
		timestamp := time.Now().Add(time.Duration(i) * time.Hour).Format("15:04")
		data[i] = TimeSeries{
			Time:  timestamp,
			Value: currentValue,
		}
	}

	delta := data[len(data)-1].Value - data[0].Value
	return data, delta
}
