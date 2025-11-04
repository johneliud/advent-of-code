package day2

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func DetermineFeetOfRibbon(input string) int {
	splitInput := strings.Split(input, "\n")

	var smallestPerimeter, volume int

	for i, inp := range splitInput {
		if len(inp) == 0 {
			continue
		}

		dimensions := strings.Split(inp, "x")
		if len(dimensions) != 3 {
			fmt.Printf("Invalid dimensions on line %v.", i+1)
			continue
		}

		length, width, height := dimensions[0], dimensions[1], dimensions[2]

		l, err := strconv.Atoi(length)
		if err != nil {
			fmt.Printf("Failed converting %v to int: %v\n", length, err)
			return 0
		}

		w, err := strconv.Atoi(width)
		if err != nil {
			fmt.Printf("Failed converting %v to int: %v\n", width, err)
			return 0
		}

		h, err := strconv.Atoi(height)
		if err != nil {
			fmt.Printf("Failed converting %v to int: %v\n", height, err)
			return 0
		}
		dimensionSlice := []int{l, w, h}

		sort.Ints(dimensionSlice)

		smallestPerimeter += dimensionSlice[0] + dimensionSlice[0] + dimensionSlice[1] + dimensionSlice[1]
		volume += l * w * h
	}
	return smallestPerimeter + volume
}
