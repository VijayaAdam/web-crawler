package main

import (
	"context"
	"go-test/crawler"
	"go-test/storage"
	"sync"
	"time"
)

func main() {
	startURL := "https://example.com/"
	numWorkers := 5

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	urlQueue := make(chan string, 200)
	visited := storage.NewVisited()
	rateLimiter := time.NewTicker(300 * time.Millisecond)

	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go crawler.Worker(
			ctx,
			i,
			urlQueue,
			visited,
			&wg,
			rateLimiter,
		)
	}

	visited.Mark(startURL)
	urlQueue <- startURL

	<-ctx.Done()
	wg.Wait()
}
