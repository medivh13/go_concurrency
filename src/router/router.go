package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go_concurrency/src/handlers"
	"go_concurrency/src/repository"
	"go_concurrency/src/usecases"
)

// NewRouter creates and returns a new router.
func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Initialize the repository
	repo := repository.NewFetcher()

	// Initialize the usecase with the repository
	uc := usecases.NewFetcherUsecase(repo)

	// Initialize the handlers with the usecase
	h := handlers.NewHandler(uc)

	// Define routes
	r.Get("/fetch-concurrent", h.ConcurrentHandler)
	r.Get("/fetch-sequential", h.SequentialHandler)

	return r
}
