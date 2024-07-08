# Go Concurrency API Example

This is an example Go project demonstrating how to use concurrency to fetch data from multiple URLs, and how to structure a Go project with multiple layers (repository, usecase, handler) using interfaces. The project also uses `go-chi` for routing.

## Project Structure

├── main.go
├── router
│ └── router.go
├── handler
│ └── handler.go
├── usecase
│ └── fetcher.go
└── repository
└── fetcher.go


## Getting Started

### Prerequisites

- Go 1.16+ installed on your machine. You can download it from [here](https://golang.org/dl/).

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/medivh13/go_concurrency.git
    cd go_concurrency
    ```

2. Initialize the Go modules:

    ```sh
    go mod tidy
    ```

### Running the Server

1. Run the server:

    ```sh
    go run main.go
    ```

    You should see the following output indicating the server is running:

    ```
    Server is running on port 8080
    ```

### Endpoints

There are two main endpoints in this project:

1. **Fetch Concurrently**: Fetches data from multiple URLs concurrently.

    - **URL**: `/fetch-concurrent`
    - **Method**: `GET`
    - **Response**: JSON object containing the fetched results and the elapsed time.

    Example response:

    ```json
    {
      "results": [
        "Fetched 1234 bytes from https://jsonplaceholder.typicode.com/posts in 0.23 seconds",
        "Fetched 5678 bytes from https://jsonplaceholder.typicode.com/comments in 0.45 seconds",
        "Fetched 91011 bytes from https://jsonplaceholder.typicode.com/albums in 0.67 seconds"
      ],
      "elapsed_time": 0.69
    }
    ```

2. **Fetch Sequentially**: Fetches data from multiple URLs sequentially.

    - **URL**: `/fetch-sequential`
    - **Method**: `GET`
    - **Response**: JSON object containing the fetched results and the elapsed time.

    Example response:

    ```json
    {
      "results": [
        "Fetched 1234 bytes from https://jsonplaceholder.typicode.com/posts in 0.23 seconds",
        "Fetched 5678 bytes from https://jsonplaceholder.typicode.com/comments in 0.45 seconds",
        "Fetched 91011 bytes from https://jsonplaceholder.typicode.com/albums in 0.67 seconds"
      ],
      "elapsed_time": 1.35
    }
    ```

### Testing with Postman

1. Open Postman.
2. Create a new GET request.
3. For concurrent fetching, set the URL to `http://localhost:8080/fetch-concurrent` and send the request.
4. For sequential fetching, set the URL to `http://localhost:8080/fetch-sequential` and send the request.

You should see the JSON response with the fetched data and the elapsed time.

## Code Overview

### main.go

This is the entry point of the application. It initializes the router and starts the server.

### router/router.go

This file sets up the routes for the API endpoints using `go-chi`.

### handler/handler.go

This file contains the HTTP handlers for the API endpoints. It uses the fetcher usecase to fetch data from URLs either concurrently or sequentially.

### usecase/fetcher.go

This file defines the usecase for fetching data from URLs. It uses a repository to fetch the actual data and provides methods for concurrent and sequential fetching.

### repository/fetcher.go

This file defines the repository for fetching data from URLs. It contains the actual implementation of the HTTP requests to fetch the data.

## Conclusion

This project demonstrates a basic yet powerful approach to structuring a Go application with multiple layers and using concurrency to improve performance. Feel free to explore and modify the code to suit your needs.

If you have any questions or suggestions, feel free to open an issue or submit a pull request.

Happy coding!
