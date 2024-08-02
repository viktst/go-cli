package fetchers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const BaseURL = "http://numbersapi.com"

func GetFacts(number interface{}, factType string) (string, error) {
	url := BuildURL(number, factType)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("http get err: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch fact: %s", resp.Status)
	}

	return ReadResponseBody(resp)
}

func BuildURL(number interface{}, factType string) string {
	if factType == "random" {
		return fmt.Sprintf("%s/random/%s", BaseURL, factType)
	}
	return fmt.Sprintf("%s/%v/%s", BaseURL, number, factType)
}

func ReadResponseBody(resp *http.Response) (string, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read body err: %w", err)
	}
	return string(body), nil
}
