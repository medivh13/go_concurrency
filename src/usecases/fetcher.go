package usecases

import "go_concurrency/src/repository"

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
    // Create a channel to communicate results from goroutines
    ch := make(chan string)
    
    // Iterate over each URL
    for _, url := range urls {
        // Launch a goroutine to fetch each URL concurrently
        go func(url string) {
            result, err := uc.fetcherRepo.FetchURL(url)
            if err != nil {
                // Send the error message to the channel if an error occurs
                ch <- err.Error()
            } else {
                // Send the result to the channel if successful
                ch <- result
            }
        }(url)
    }

    var results []string
    // Collect results from the channel for each URL
    for range urls {
        results = append(results, <-ch)
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

