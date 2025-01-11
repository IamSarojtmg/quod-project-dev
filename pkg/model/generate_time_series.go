package model

import (
	"math/rand"
	"time"
)

type TimeSeries struct {
	Time  string
	Value int
}

func GenerateTimeSeries() ([]TimeSeries, int) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	data := make([]TimeSeries, 10)

	startValue := 50 + rng.Intn(20)
	currentValue := startValue

	for i := 0; i < 10; i++ {
		trend := 1 - rng.Intn(3)
		noise := rng.Intn(3) - 1

		currentValue += trend + noise
		if currentValue < 0 {
			currentValue = 0
		}

		timestamp := time.Now().Add(time.Duration(i) * time.Hour).Format("15:04")
		data[i] = TimeSeries{
			Time:  timestamp,
			Value: currentValue,
		}
	}

	delta := data[len(data)-1].Value - data[0].Value
	return data, delta
}
