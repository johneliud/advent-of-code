package main

import (
	"fmt"
	"strconv"
	"strings"
)

func determineBlocksDistance(input string) int {
	var xPos, yPos, currentDirection int
	directions := []string{"N", "E", "S", "W"}

	trimmedInput := strings.Fields(input)

	for _, inp := range trimmedInput {
		direction := string(inp[0])
		distance := string(inp[1:])

		if len(distance) >= 2 {
			distance = distance[0 : len(distance)-1]
		}

		blocks, err := strconv.Atoi(distance)
		if err != nil {
			fmt.Printf("Failed converting %v to int: %v\n", distance, err)
			return 0
		}

		if strings.ToUpper(direction) == "R" {
			currentDirection = (currentDirection + 1) % 4
		} else if strings.ToUpper(direction) == "L" {
			currentDirection = (currentDirection + 3) % 4
		}

		switch directions[currentDirection] {
		case "N":
			yPos += blocks
		case "E":
			xPos += blocks
		case "S":
			yPos -= blocks
		case "W":
			xPos -= blocks
		}
	}
	blocksDistance := abs(xPos) + abs(yPos)
	return blocksDistance
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
