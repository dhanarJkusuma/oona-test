package usecase_test

import (
	"oona-test/usecase"
	"testing"
)

func TestFindHighSum(t *testing.T) {
	expect := []struct {
		input    []float64
		expected float64
	}{
		{[]float64{3, -9, -1, 4, 3, -2, 3}, 8},
	}

	for _, val := range expect {
		if res := usecase.GetHighest(val.input); res != val.expected {
			t.Error("Test Calculate Failed : {} inputer, {} expected, and received {}", val.input, val.expected, res)
		}

	}
}
