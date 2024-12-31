package main

func findPosition(s string) int {
	var floor, position int

	for i, char := range s {
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		} else {
			continue
		}

		if floor == -1 {
			position = i + 1
			break
		}
	}
	return position
}
