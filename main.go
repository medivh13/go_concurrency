package main

import (
	"fmt"
	"log"
	"net/http"

	"go_concurrency/src/router"
)

func main() {
    // Initialize the router
    r := router.NewRouter()

    // Print a message indicating that the server is running
    fmt.Println("Server is running on port 8080")

    // Start the HTTP server on port 8080 and use the router for handling requests
    log.Fatal(http.ListenAndServe(":8080", r))
}

