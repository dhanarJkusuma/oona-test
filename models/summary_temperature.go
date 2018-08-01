package models

// SummaryTemperature represent fridge sensor data summary
type SummaryTemperature struct {
	ID      string    `json:"id"`
	Average string    `json:"average"`
	Median  string    `json:"median"`
	Mode    []float64 `json:"mode"`
}
