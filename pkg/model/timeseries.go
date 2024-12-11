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
