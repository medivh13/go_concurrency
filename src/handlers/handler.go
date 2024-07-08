package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"go_concurrency/src/usecases"
)

type Handler struct {
	fetcherUsecase usecases.FetcherUsecase
}

func NewHandler(uc usecases.FetcherUsecase) *Handler {
	return &Handler{fetcherUsecase: uc}
}

// ConcurrentHandler handles the /fetch-concurrent endpoint.
func (h *Handler) ConcurrentHandler(w http.ResponseWriter, r *http.Request) {
	urls := []string{
		"https://jsonplaceholder.typicode.com/posts",
		"https://jsonplaceholder.typicode.com/comments",
		"https://jsonplaceholder.typicode.com/albums",
	}
	start := time.Now()
	results := h.fetcherUsecase.FetchURLsConcurrently(urls)
	duration := time.Since(start).Seconds()

	response := struct {
		Results     []string `json:"results"`
		ElapsedTime float64  `json:"elapsed_time"`
	}{
		Results:     results,
		ElapsedTime: duration,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// SequentialHandler handles the /fetch-sequential endpoint.
func (h *Handler) SequentialHandler(w http.ResponseWriter, r *http.Request) {
	urls := []string{
		"https://jsonplaceholder.typicode.com/posts",
		"https://jsonplaceholder.typicode.com/comments",
		"https://jsonplaceholder.typicode.com/albums",
	}
	start := time.Now()
	results := h.fetcherUsecase.FetchURLsSequentially(urls)
	duration := time.Since(start).Seconds()

	response := struct {
		Results     []string `json:"results"`
		ElapsedTime float64  `json:"elapsed_time"`
	}{
		Results:     results,
		ElapsedTime: duration,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
