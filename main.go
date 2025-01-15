package main

import (
	"fmt"
	"os"

	"advent-of-code/2015/day1"
	"advent-of-code/utils"

	"github.com/joho/godotenv"
)

var year, day string

func main() {
	switch len(os.Args) {
	case 3:
		year = os.Args[1]
		day = os.Args[2]
	default:
		fmt.Println("Usage: 'go run . [YEAR] [DAY]' OR 'make [YEAR] [DAY]'")
		return
	}

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Failed loading .env file: %v\n", err)
		return
	}

	input := utils.FetchInput(year, day)

	fmt.Printf(`Day %v

Part 1: %v
Part 2: %v
`, day, day1.FindFloor(input), day1.FindPosition(input))
}
