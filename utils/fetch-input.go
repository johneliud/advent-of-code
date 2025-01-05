package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const baseURL = "https://adventofcode.com/%s/day/%s/input"

// FetchInput accepts year and day as parameters in order to fetch the puzzle input from the specified URL.
func FetchInput(year, day string) string {
	url := fmt.Sprintf(baseURL, year, day)

	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Failed to make a request: %v\n", err)
		return ""
	}

	req.Header.Add("Cookie", "session="+os.Getenv("AOC_SESSION_COOKIE"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Failed to get a response: %v\n", err)
		return ""
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected status code: %d\n", res.StatusCode)
		return ""
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Failed reading response body: %v\n", err)
		return ""
	}
	return string(body)
}
