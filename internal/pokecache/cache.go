package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Data map[string]cacheEntry
	mu   *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {

	return Cache
}

func (cache *Cache) Add(key string, val []byte) {

	cache.mu.Lock()
	defer cache.mu.Unlock()

	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	cache.Data[key] = newEntry
}

func (cache *Cache) Get(key string) ([]byte, bool) {

	cache.mu.RLock()
	defer cache.mu.RUnlock()

	requestedData, ok := cache.Data[key]
	if !ok {
		return []byte{}, false
	}

	return requestedData.val, true
}

func (cache *Cache) reaploop() {

}
