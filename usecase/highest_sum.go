package usecase

import (
	m "math"
)

func GetHighest(input []float64) float64 {
	var meh, msf float64
	meh = input[0]
	msf = input[0]
	for i, val := range input {
		if i >= 1 {
			meh = m.Max(val, meh+val)
			msf = m.Max(msf, meh)
		}
	}
	return msf
}
