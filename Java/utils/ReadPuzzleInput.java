package utils;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;

public class ReadPuzzleInput {
    public static String[] readPuzzleInput(String filename) {
        try {
            return Files.readAllLines(Paths.get(filename)).toArray(new String[0]);
        } catch (IOException e) {
            System.out.println("Error reading puzzle file: " + filename);
            throw new RuntimeException(e);
        }
    }
}
