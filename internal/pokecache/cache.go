package pokecache

import (
	"time"
)

type Cache struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	return Cache
}
