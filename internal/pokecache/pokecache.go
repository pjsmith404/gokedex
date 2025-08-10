package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu *sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		entries: make(map[string]cacheEntry),
		mu: &sync.Mutex{},
		interval: interval,
	}

	go c.reapLoop()

	return c
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c Cache) Get(key string) (val []byte, found bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	val = c.entries[key].val

	if val == nil {
		return []byte{}, false
	}

	return val, true
}

func (c Cache) reapLoop() {
	for t := range time.Tick(c.interval) {
		for k, v := range c.entries {
			age := t.Sub(v.createdAt)
			if age > c.interval {
				c.mu.Lock()
				delete(c.entries, k)
				c.mu.Unlock()
			}
		}
	}
}

