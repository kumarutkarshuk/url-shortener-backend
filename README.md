# URL Shortener

This is a simple URL shortener service built with Go. It uses Redis for storing the shortened URLs and Gorilla Mux for routing.


## Setup

1. **Clone the repository:**
    ```sh
    git clone <repository-url>
    cd url-shortener
    ```

2. **Install dependencies:**
    ```sh
    go mod tidy
    ```

3. **Run the server:**
    ```sh
    go run main.go
    ```

## Endpoints

- **POST /shorten**
    - Request: `{"URL": "http://example.com"}`
    - Response: `{"short_url": "http://localhost:8080/<short-URL-hash>"}`

- **GET /{shortURL}**
    - Redirects to the original URL.