package _2015.day_one;

public class PartTwo {
    public static int findPosition(String input) {
        int floor = 0, position = 0;
        
        for (char c : input.toCharArray()) {
            position++;
            if (c == '(') {
                floor++;
            } else if (c == ')') {
                floor--;
            } else {
                continue;
            }
            
            if (floor == -1) {
                break;
            }
        }
        return position;
    }
}