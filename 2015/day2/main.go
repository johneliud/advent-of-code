package main

import (
	"advent-of-code/utils"
	"fmt"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	start := time.Now()

	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Printf("Failed loading .env file: %v\n", err)
		return
	}

	input := utils.FetchInput("2015", "2")

	fmt.Println("Wrapping Paper:", determineSquareFeetOfWrappingPaper(input))
	fmt.Println("Ribbon:", determineFeetOfRibbon(input))

	elapsed := time.Since(start)
	fmt.Println("\nTime taken:", elapsed)
}
