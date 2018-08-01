package usecase_test

import (
	"oona-test/models"
	"oona-test/usecase"
	"testing"
)

func TestTemperature(t *testing.T) {
	expect := []struct {
		input           []models.PayloadTemperature
		expected_id     string
		expected_avg    string
		expected_median string
		expected_mode   []float64
	}{
		{
			[]models.PayloadTemperature{
				{"c", 1509493643, 3.96},
				{"c", 1509493645, 3.96},
				{"c", 1509493655, 3.95},
				{"c", 1510127886, 3.36},
				{"c", 1510127892, 3.36},
			},
			"c",
			"3.72",
			"3.96",
			[]float64{3.36, 3.96},
		},
	}

	for _, val := range expect {
		res := usecase.GetSummary(val.input)
		summary := res[0]

		if summary.ID != val.expected_id {
			t.Error("Test Summary Failed : ID {} expected, and received {}", val.expected_id, summary.ID)
		}

		if summary.Average != val.expected_avg {
			t.Error("Test Summary Failed : Average {} expected, and received {}", val.expected_avg, summary.Average)
		}

		if summary.Median != val.expected_median {
			t.Error("Test Summary Failed : Median {} expected, and received {}", val.expected_median, summary.Median)
		}

		for i, _ := range summary.Mode {
			if summary.Mode[i] != val.expected_mode[i] {
				t.Error("Test Summary Failed : Mode {} expected, and received {}", val.expected_mode, summary.Mode)
			}
		}

	}
}
