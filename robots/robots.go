package robots

import (
	"net/http"
	"net/url"
	"sync"

	"github.com/temoto/robotstxt"
)

var (
	mu    sync.Mutex
	cache = make(map[string]*robotstxt.RobotsData)
)

// Allowed returns true if the given URL is allowed to be crawled
// according to the site's robots.txt rules.
func Allowed(target string) bool {
	u, err := url.Parse(target)
	if err != nil {
		return false
	}

	// robots.txt is defined per scheme + host
	host := u.Scheme + "://" + u.Host

	// Try cache first
	mu.Lock()
	data, ok := cache[host]
	mu.Unlock()

	// Fetch robots.txt only once per host
	if !ok {
		resp, err := http.Get(host + "/robots.txt")
		if err != nil {
			// If robots.txt can't be fetched, default to allow
			return true
		}
		defer resp.Body.Close()

		data, err = robotstxt.FromResponse(resp)
		if err != nil {
			return true
		}

		// Save to cache
		mu.Lock()
		cache[host] = data
		mu.Unlock()
	}

	// Use cached rules
	group := data.FindGroup("*")
	return group.Test(u.Path)
}
