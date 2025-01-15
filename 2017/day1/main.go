package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"advent-of-code/utils"

	"github.com/joho/godotenv"
)

func main() {
	start := time.Now()

	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Printf("Failed loading .env file: %v\n", err)
		return
	}

	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed getting working directory: %v\n", err)
		return
	}

	splitWd := strings.Split(wd, "/")
	year := splitWd[len(splitWd)-2]
	day := splitWd[len(splitWd)-1][3:]

	input := utils.FetchInput(year, day)

	fmt.Println(input)

	elapsed := time.Since(start)
	fmt.Println("\nTime taken:", elapsed)
}
