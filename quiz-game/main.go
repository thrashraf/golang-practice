package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file := "./problem.csv"
	var result []bool
	var correct int

	csvFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
	}

	result = calculateResult(records)

	for _, res := range result {
		if res {
			correct++
		}
	}

	fmt.Println("You scored", correct, "out of", len(records))
}

func calculateResult(records [][]string) []bool {
	var result []bool

	for _, record := range records {
		question := record[0]
		answer := record[1]

		var sum int

		for _, char := range question {
			if char != '+' {
				sum += int(char - '0')
			}
		}

		if fmt.Sprint(sum) == answer {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}

	return result
}
