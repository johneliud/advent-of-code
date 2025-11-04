package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const baseURL = "https://adventofcode.com/%s/day/%s/input"

// FetchInput accepts year and day as parameters in order to fetch the puzzle input from the specified URL.
func FetchInput(year, day string) (string, error) {
	url := fmt.Sprintf(baseURL, year, day)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to make a request: %v", err)
	}

	session := os.Getenv("AOC_SESSION_COOKIE")
	if len(session) == 0 {
		return "", fmt.Errorf("environment varible not set")
	}

	req.Header.Add("Cookie", "session="+session)

	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to get a response: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed reading response body: %v", err)
	}
	return string(body), nil
}
