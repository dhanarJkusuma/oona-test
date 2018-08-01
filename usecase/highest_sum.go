package usecase

import "math"

func GetHighest(input []float64) float64 {
	var meh, msf float64
	meh = input[0]
	msf = input[0]
	for i, val := range input {
		if i >= 1 {
			meh = math.Max(val, meh+val)
			msf = math.Max(msf, meh)
		}
	}
	return msf
}
