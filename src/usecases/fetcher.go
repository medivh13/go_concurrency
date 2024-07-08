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
	ch := make(chan string)
	for _, url := range urls {
		go func(url string) {
			result, err := uc.fetcherRepo.FetchURL(url)
			if err != nil {
				ch <- err.Error()
			} else {
				ch <- result
			}
		}(url)
	}

	var results []string
	for range urls {
		results = append(results, <-ch)
	}

	return results
}

// FetchURLsSequentially fetches multiple URLs sequentially.
func (uc *fetcherUsecase) FetchURLsSequentially(urls []string) []string {
	var results []string
	for _, url := range urls {
		result, err := uc.fetcherRepo.FetchURL(url)
		if err != nil {
			results = append(results, err.Error())
		} else {
			results = append(results, result)
		}
	}
	return results
}
