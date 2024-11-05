package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {

	c := Cache{
		cache: make(map[string]cacheEntry),
		mu:    new(sync.RWMutex),
	}

	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}

}

func (c *Cache) Get(key string) ([]byte, bool) {

	c.mu.RLock()
	defer c.mu.RUnlock()

	requestedEntry, ok := c.cache[key]
	if !ok {
		return []byte{}, false
	}

	return requestedEntry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {

	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	timeAgo := time.Now().UTC().Add(-interval)

	c.mu.Lock()
	defer c.mu.Unlock()

	for i, entry := range c.cache {
		if entry.createdAt.Before(timeAgo) {
			delete(c.cache, i)
		}
	}
}
