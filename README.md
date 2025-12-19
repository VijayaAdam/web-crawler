# Go Concurrent Web Crawler

Designed and implemented a concurrent, robots-compliant web crawler in Go using goroutines and channels. The crawler efficiently discovers and processes web pages with proper rate limiting, context-based cancellation and thread-safe URL deduplication, while respecting crawl policies through per-domain caching.

---

## Features

- Concurrent web crawling using goroutines and worker pool
- Context-based cancellation with timeout for graceful shutdown
- Robots.txt compliance with per-domain caching
- Rate-limited HTTP requests to ensure polite crawling
- Thread-safe URL deduplication using mutex-protected storage
- HTML parsing and link extraction with relative URL resolution
- Non-blocking URL queue with backpressure handling
- Modular architecture with clear separation of concerns

---

## 🗂️ Project Structure

    .
    ├── crawler/          # Worker logic and HTTP crawling
    │   └── worker.go
    ├── parser/           # HTML parsing and link extraction
    │   └── parser.go
    ├── robots/           # robots.txt fetching, parsing and caching
    │   └── robots.go
    ├── storage/          # Thread-safe visited URL storage
    │   └── storage.go
    ├── main.go           # Application entry point
    ├── go.mod
    └── go.sum


## How It Works

1. Starts crawling from a seed URL
2. Multiple workers fetch pages concurrently
3. Extracts links from HTML responses
4. Checks robots.txt rules before enqueueing URLs
5. Deduplicates URLs to prevent re-crawling
6. Uses rate limiting to avoid overwhelming servers
7. Stops gracefully after timeout

---

## Running the Project

### Prerequisites
- Go 1.20+

### Run
```bash
go mod tidy
go run main.go
