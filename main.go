package main

import (
	"fmt"

	"advent-of-code/2015/day1"
	"advent-of-code/utils"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Failed loading .env file: %v\n", err)
		return
	}

	input := utils.FetchInput("2015", "1")

	fmt.Printf(`Day 1

Part 1: %v
Part 2: %v
`, day1.FindFloor(input), day1.FindPosition(input))
}
