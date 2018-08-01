package models

// PayloadTemperature represent each data in fridge sensor
type PayloadTemperature struct {
	ID          string  `json:"id"`
	Timestamp   int64   `json:"timestamp"`
	Temperature float64 `json:"temperature"`
}
