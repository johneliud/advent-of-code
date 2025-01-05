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

	input := utils.FetchInput("2015", "1")

	fmt.Println("Floor:", findFloor(input))
	fmt.Println("Position:", findPosition(input))

	elapsed := time.Since(start)
	fmt.Println("\nTime taken:", elapsed)
}
