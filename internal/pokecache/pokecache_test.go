package pokecache

import (
	"testing"
	"time"
)

func TestReap(t *testing.T) {
	const interval = 5 * time.Second

	cache := NewCache(interval)

	cache.Add("key1", []byte("val1"))

	time.Sleep(interval + time.Second)

	_, ok := cache.Get("key1")
	if ok {
		t.Error("old entry was not reaped")
	}
}
