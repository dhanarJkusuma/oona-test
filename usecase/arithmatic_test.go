package usecase_test

import (
	"oona-test/usecase"
	"testing"
)

func TestCalculate(t *testing.T) {
	expect := []struct {
		input    string
		expected string
	}{
		{"satu tambah satu", "dua"},
		{"dua belas kurang satu", "sebelas"},
		{"satu dua tambah satu", "tiga belas"},
	}

	for _, val := range expect {
		if res := usecase.Calculate(val.input); res != val.expected {
			t.Error("Test Calculate Failed : {} inputer, {} expected, and received {}", val.input, val.expected, res)
		}

	}
}
