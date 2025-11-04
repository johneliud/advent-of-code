package day1

func FindFloor(s string) int {
	var floor int

	for _, char := range s {
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}
	}
	return floor
}
