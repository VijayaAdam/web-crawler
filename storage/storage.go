package storage

import "sync"

type Visited struct {
	mu   sync.Mutex
	urls map[string]bool
}

func (v *Visited) CheckAndMark(link string) bool {
	v.mu.Lock()
	defer v.mu.Unlock()

	if v.urls[link] {
		return true
	}
	v.urls[link] = true
	return false
}

func NewVisited() *Visited {
	return &Visited{
		urls: make(map[string]bool),
	}
}

func (v *Visited) Mark(u string) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.urls[u] = true
}
