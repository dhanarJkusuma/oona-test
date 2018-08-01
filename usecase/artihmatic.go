package usecase

import (
	"idr-cardinalparser/converter"
	"log"
	"strings"
)

func Calculate(input string) string {
	var cardinal []string
	var operation string
	if ops, found := IsContainArithmaticOperation(input); found {
		cardinal = strings.Split(input, ops)
		operation = ops
	}
	first := cardinal[0]
	second := cardinal[1]

	conv := converter.NewNumberConverter()
	num1, err := conv.ConvertToNumber(first)
	if err != nil {
		log.Fatalln("Invalid input")
	}

	num2, err := conv.ConvertToNumber(second)
	if err != nil {
		log.Fatalln("Invalid input")
	}

	var result int64
	result = 0
	switch operation {
	case "tambah":
		result = num1 + num2
		break
	case "kurang":
		result = num1 - num2
		break
	case "kali":
		result = num1 * num2
		break
	case "bagi":
		result = num1 / num2
		break
	}

	rs, err := conv.ConvertToWords(result)
	if err != nil {
		log.Fatalln("Invalid result")
	}
	return rs
}

func IsContainArithmaticOperation(input string) (string, bool) {
	op := []string{"tambah", "kurang", "kali", "bagi"}
	for _, val := range op {
		if strings.Contains(input, val) {
			return val, true
		}
	}
	return "", false
}
