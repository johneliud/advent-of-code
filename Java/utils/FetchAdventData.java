package utils;

import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.stream.Stream;

public class FetchAdventData {
    /*
     * Fetch the advent data from the server instead of copying the data to a local file.
     */
    private static final String BASE_URL = "https://adventofcode.com/%d/day/%d/input";

    public static String fetchAdventData(String year, String day) {
        try {
            int yearInt = Integer.parseInt(year);
            int dayInt = Integer.parseInt(day);
            String url = String.format(BASE_URL, yearInt, dayInt);

            String sessionCookie = System.getenv("AOC_SESSION_COOKIE");

            // Try reading from .env file
            if (sessionCookie == null || sessionCookie.isEmpty()) {
                try (Stream<String> lines = Files.lines(Path.of(".env"))) {
                    sessionCookie = lines
                            .filter(line -> line.startsWith("AOC_SESSION_COOKIE="))
                            .map(line -> line.split("=", 2)[1].trim())
                            .findFirst()
                            .orElse(null);
                } catch (IOException e) {
                    System.err.println("Could not read .env file: " + e.getMessage());
                }
            }

            HttpClient client = HttpClient.newHttpClient();
            HttpRequest.Builder requestBuilder = HttpRequest.newBuilder()
                    .uri(URI.create(url))
                    .GET();

            if (sessionCookie != null && !sessionCookie.isEmpty()) {
                requestBuilder.header("Cookie", "session=" + sessionCookie);
            }

            HttpRequest request = requestBuilder.build();
            HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());

            if (response.statusCode() == 200) {
                return response.body();
            } else {
                System.err.println("Failed to fetch data. Status code: " + response.statusCode());
                return null;
            }
        } catch (IOException | InterruptedException e) {
            System.err.println("Error fetching advent data: " + e.getMessage());
            return null;
        } catch (NumberFormatException e) {
            System.err.println("Invalid year or day format: " + e.getMessage());
            return null;
        }
    }
}
