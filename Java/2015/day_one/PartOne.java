package _2015.day_one;

public class PartOne {
    public static int findFloor(String input) {
        int floor = 0;

        for (char c : input.toCharArray()) {
            if (c == '(') {
                floor++;
            } else if (c == ')') {
                floor--;
            }
        }
        return floor;
    }
}
