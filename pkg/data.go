package pkg

import (
	"math/rand"
	"time"
)

var SensorTypes = []string{"a", "b", "c", "d", "e", "f", "g"}
var SensorLocations = []string{"bottom", "top", "left", "right", "middle", "under", "above"}

func GenerateData(amount int) (data []InsertSchema) {
	for i := 0; i <= amount; i++ {
		data = append(data, InsertSchema{
			Time:        time.Now(),
			SensorId:    rand.Intn(len(SensorTypes)) + 1, // +1 because sensorId can't be 0
			Temperature: rand.Float64() * float64(rand.Intn(99)),
			Cpu:         rand.Float64() * float64(rand.Intn(99)),
		})
	}

	return data
}
