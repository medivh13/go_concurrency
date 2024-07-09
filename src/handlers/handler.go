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
	// List of URLs to be fetched concurrently
	urls := []string{
		"https://jsonplaceholder.typicode.com/posts",
		"https://jsonplaceholder.typicode.com/comments",
		"https://jsonplaceholder.typicode.com/albums",
	}
	// Start recording the time
	start := time.Now()
	// Fetch URLs concurrently
	results := h.fetcherUsecase.FetchURLsConcurrently(urls)
	// Calculate the elapsed time
	duration := time.Since(start).Seconds()

	// Structure the response with results and elapsed time
	response := struct {
		Results     []string `json:"results"`
		ElapsedTime float64  `json:"elapsed_time"`
	}{
		Results:     results,
		ElapsedTime: duration,
	}

	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")
	// Encode and send the response as JSON
	json.NewEncoder(w).Encode(response)
}

// SequentialHandler handles the /fetch-sequential endpoint.
func (h *Handler) SequentialHandler(w http.ResponseWriter, r *http.Request) {
	// List of URLs to be fetched sequentially
	urls := []string{
		"https://jsonplaceholder.typicode.com/posts",
		"https://jsonplaceholder.typicode.com/comments",
		"https://jsonplaceholder.typicode.com/albums",
	}
	// Start recording the time
	start := time.Now()
	// Fetch URLs sequentially
	results := h.fetcherUsecase.FetchURLsSequentially(urls)
	// Calculate the elapsed time
	duration := time.Since(start).Seconds()

	// Structure the response with results and elapsed time
	response := struct {
		Results     []string `json:"results"`
		ElapsedTime float64  `json:"elapsed_time"`
	}{
		Results:     results,
		ElapsedTime: duration,
	}

	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")
	// Encode and send the response as JSON
	json.NewEncoder(w).Encode(response)
}
