package _2025.DayOne;

import utils.ReadPuzzleInput;

public class PartOne {
	public static int findPassword(String[] puzzleInput) {
		int password = 0;
		int currentDialPosition = 50;

		for (String line : puzzleInput) {
			String direction = String.valueOf(line.charAt(0));
			int value = Integer.parseInt(line.substring(1));

			if (direction.equals("R")) {
				currentDialPosition += value;
			} else if (direction.equals("L")) {
				currentDialPosition -= value;
			} else {
				continue;
			}

			currentDialPosition = (currentDialPosition % 100 + 100) % 100;

			if (currentDialPosition == 0) {
				password++;
			}
		}

		return password;
	}
}
