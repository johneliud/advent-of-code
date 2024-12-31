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

	input := utils.ReadInput("2015", "2")

	fmt.Println("Square feet of wrapping paper:", determineSquareFeetOfWrappingPaper(input))
	fmt.Println("Feet of ribbon:", determineFeetOfRibbon(input))

	elapsed := time.Since(start)
	fmt.Println("\nTime taken:", elapsed)
}
