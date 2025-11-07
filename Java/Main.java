import utils.FetchAdventData;
import _2015.day_one.PartOne;

public class Main {
    public static void main(String[] args) {
        if (args.length < 2) {
            System.out.println("Usage: java -cp build Main <year> <day>");
            return;
        }

        String year = args[0];
        String day = args[1];

        String adventData = FetchAdventData.fetchAdventData(year, day);

        if (adventData != null) {
            System.out.println("Advent of code data fetched successfully.\n\n");

            int floor = PartOne.findFloor(adventData);

            System.out.println("Floor: " + floor);
        } else {
            System.out.println("Could not fetch advent data for year " + year + ", day " + day + ".");
        }
    }
}
