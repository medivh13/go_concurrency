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
    // Start recording the time when this function is executed
    start := time.Now()
    
    // Send an HTTP GET request to the given URL
    resp, err := http.Get(url)
    if err != nil {
        // If there is an error while sending the request, return the error
        return "", fmt.Errorf("error fetching %s: %v", url, err)
    }
    // Close resp.Body after the function is finished executing
    defer resp.Body.Close()

    // Read the entire response body from the server
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // If there is an error while reading the response body, return the error
        return "", fmt.Errorf("error reading response body from %s: %v", url, err)
    }

    // Calculate the duration it took to fetch the data from the URL
    duration := time.Since(start).Seconds()

    // Return the number of bytes fetched, the URL, and the duration it took
    return fmt.Sprintf("Fetched %d bytes from %s in %.2f seconds", len(body), url, duration), nil
}


