package repository

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// FetcherRepository defines the interface for the fetcher repository.
type FetcherRepository interface {
	FetchURL(url string) (string, error)
}

type fetcherRepo struct{}

// NewFetcher creates a new instance of fetcher repository.
func NewFetcher() FetcherRepository {
	return &fetcherRepo{}
}

// FetchURL is a helper function that fetches the content of a URL.
func (r *fetcherRepo) FetchURL(url string) (string, error) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error fetching %s: %v", url, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body from %s: %v", url, err)
	}

	duration := time.Since(start).Seconds()
	return fmt.Sprintf("Fetched %d bytes from %s in %.2f seconds", len(body), url, duration), nil
}
