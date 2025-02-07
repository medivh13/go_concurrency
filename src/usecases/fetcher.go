package usecases

import (
	"go_concurrency/src/repository"
	"sync"
)

// FetcherUsecase defines the interface for fetcher usecases.
type FetcherUsecase interface {
	FetchURLsConcurrently(urls []string) []string
	FetchURLsSequentially(urls []string) []string
}

type fetcherUsecase struct {
	fetcherRepo repository.FetcherRepository
}

// NewFetcherUsecase creates a new instance of fetcher usecase.
func NewFetcherUsecase(repo repository.FetcherRepository) FetcherUsecase {
	return &fetcherUsecase{fetcherRepo: repo}
}

// FetchURLsConcurrently fetches multiple URLs concurrently.
func (uc *fetcherUsecase) FetchURLsConcurrently(urls []string) []string {
	var wg sync.WaitGroup              // WaitGroup to wait for all goroutines to complete
	ch := make(chan string, len(urls)) // Buffered channel to prevent blocking

	// Iterate over each URL and launch a goroutine
	for _, url := range urls {
		wg.Add(1) // Increment the WaitGroup counter
		go func(url string) {
			defer wg.Done() // Decrement the counter when the goroutine finishes
			result, err := uc.fetcherRepo.FetchURL(url)
			if err != nil {
				ch <- err.Error() // Send the error message to the channel
			} else {
				ch <- result // Send the successful result to the channel
			}
		}(url)
	}

	// Launch a separate goroutine to close the channel when all goroutines finish
	go func() {
		wg.Wait()
		close(ch)
	}()

	var results []string
	// Collect results from the channel
	for result := range ch {
		results = append(results, result)
	}

	return results
}

// FetchURLsSequentially fetches multiple URLs sequentially.
func (uc *fetcherUsecase) FetchURLsSequentially(urls []string) []string {
	var results []string
	// Iterate over each URL and fetch them one by one
	for _, url := range urls {
		result, err := uc.fetcherRepo.FetchURL(url)
		if err != nil {
			// Append the error message to results if an error occurs
			results = append(results, err.Error())
		} else {
			// Append the result to results if successful
			results = append(results, result)
		}
	}
	return results
}
