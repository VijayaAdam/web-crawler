package crawler

import (
	"context"
	"fmt"
	"go-test/parser"
	"go-test/robots"
	"go-test/storage"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/html"
)

func Worker(
	ctx context.Context,
	id int,
	urls chan string,
	visited *storage.Visited,
	wg *sync.WaitGroup,
	rateLimiter *time.Ticker,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return

		case link, ok := <-urls:
			if !ok {
				return
			}

			<-rateLimiter.C

			fmt.Printf("[Worker %d] Crawling %s\n", id, link)

			req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
			if err != nil {
				continue
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				continue
			}

			if resp.StatusCode != http.StatusOK {
				resp.Body.Close()
				continue
			}

			doc, err := html.Parse(resp.Body)
			resp.Body.Close()
			if err != nil {
				continue
			}

			found := parser.ExtractLinks(link, doc)
			for _, u := range found {

				// Robots check BEFORE enqueue (best practice)
				if !robots.Allowed(u) {
					fmt.Println("Blocked by robots.txt:", u)
					continue
				}

				// Deduplication
				if visited.CheckAndMark(u) {
					continue
				}

				// Non-blocking enqueue (prevents deadlocks)
				select {
				case urls <- u:
				case <-ctx.Done():
					return
				default:
					// queue full → drop URL politely
				}
			}
		}
	}
}
