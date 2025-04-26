package pokecache

import (
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt time.Time 
	val       []byte 
}

type Cache struct {
	entries   map[string]cacheEntry
	interval  time.Duration
	mu        *sync.Mutex
}


func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
		mu:       &sync.Mutex{},
	}

	cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:	   val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.entries[key]
	if exists {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)

	go func() {
		for range ticker.C {
			c.reap()
		}
	} ()
}

func (c *Cache) reap() {
	now := time.Now()
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, entry := range c.entries {
		age := now.Sub(entry.createdAt)

		if age > c.interval {
			delete(c.entries, key)
		}
	}
}
