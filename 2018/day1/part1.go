package day1

import (
	"fmt"
	"strconv"
	"strings"
)

func DetermineFrequency(s string) int {
	var result int
	trimmedInput := strings.Fields(s)

	for _, input := range trimmedInput {
		if len(input) < 2 {
			fmt.Println("Invalid input format.")
			return 0
		}

		sign := rune(input[0])
		digit := input[1:]

		num, err := strconv.Atoi(digit)
		if err != nil {
			fmt.Printf("Failed converting %v to int: %v\n", digit, err)
			return 0
		}

		switch sign {
		case '+':
			result += num
		case '-':
			result -= num
		default:
			continue
		}
	}
	return result
}
