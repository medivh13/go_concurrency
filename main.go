package main

import (
	"fmt"
	"log"
	"net/http"

	"go_concurrency/src/router"
)

func main() {
	r := router.NewRouter()
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
