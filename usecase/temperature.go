package usecase

import (
	"fmt"
	"math"
	"oona-test/models"
	"sort"
)

func GetSummary(data []models.PayloadTemperature) []models.SummaryTemperature {
	var keyTemperature map[string][]models.PayloadTemperature
	var summaries []models.SummaryTemperature

	keyTemperature = make(map[string][]models.PayloadTemperature)

	// populate all data depend on their ID
	for _, val := range data {
		d := keyTemperature[val.ID]
		if d != nil {
			d = append(d, val)
			keyTemperature[val.ID] = d
		} else {
			ar := []models.PayloadTemperature{val}
			keyTemperature[val.ID] = ar
		}
	}

	// calculate mean
	for key, val := range keyTemperature {
		var tmp float64
		var ml map[float64]int
		var mel []float64
		var median string
		tmp = 0
		ml = make(map[float64]int)
		count := len(val)

		for _, t := range val {
			tmp += t.Temperature
			mel = append(mel, t.Temperature)
			if ml[t.Temperature] != 0 {
				ml[t.Temperature] = ml[t.Temperature] + 1
			} else {
				ml[t.Temperature] = 1
			}
		}

		sort.Float64s(mel)

		mid := count / 2
		if count%2 == 0 {
			med := ((mel[mid] + mel[mid+1]) / 2)
			median = fmt.Sprintf("%v", math.Round(med*100)/100)
		} else {
			med := mel[mid+1]
			median = fmt.Sprintf("%v", math.Round(med*100)/100)
		}

		mode := getKeyByMaxValue(ml)
		m := float64(tmp / float64(count))
		mean := fmt.Sprintf("%v", math.Ceil(m*100)/100)

		summ := models.SummaryTemperature{
			ID:      key,
			Average: mean,
			Median:  median,
			Mode:    mode,
		}
		summaries = append(summaries, summ)
	}
	return summaries
}

func getKeyByMaxValue(data map[float64]int) []float64 {
	mv := 0
	var mk []float64
	for key, value := range data {
		if value > mv {
			mv = value
			mk = mk[:0]
			mk = append(mk, key)
		} else if value == mv {
			mk = append(mk, key)
		}
	}
	return mk
}
